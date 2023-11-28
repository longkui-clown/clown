package core

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/longkui-clown/clown/pkg/utils"
)

type Manager struct {
	name        string
	exitChan    chan interface{}
	sigs        []os.Signal
	killWaitTTL time.Duration
	services    []Service
	app         Service
	logger      Logger

	// 起服钩子函数
	beforeStartup       Option
	afterServiceStartup Option
	afterAppStartup     Option

	// 关服钩子函数
	beforeShutdown       Option
	afterServiceShutdown Option
	afterAppShutdown     Option
}

func NewManager(opts ...Option) *Manager {
	manager := &Manager{
		exitChan: make(chan interface{}),
	}

	for _, opt := range opts {
		opt(manager)
	}

	if manager.logger == nil {
		manager.logger = new(DefaultManagerLogger)
	}

	return manager
}

func (m *Manager) GetApp() Service {
	return m.app
}

func (m *Manager) Logger() Logger {
	return m.logger
}

func (m *Manager) Run() {
	m.logger.INFO("------------------- STARTING -------------------")
	m.callback(m.beforeStartup)

	// 前置服务启动
	for _, service := range m.services {
		m.startService(service)
	}
	m.callback(m.afterServiceStartup)

	m.logger.INFO("App starting ...")
	m.startService(m.app)
	m.callback(m.afterAppStartup)

	m.logger.INFO("------------------- STARTED -------------------")
	m.logger.INFO("-----------------------------------------------")
	m.checkSignal()
	m.logger.INFO("-----------------------------------------------")
	m.logger.INFO("------------------- STOPPING... ---------------")
	m.callback(m.beforeShutdown)

	m.stopService(m.app)
	m.callback(m.afterAppShutdown)

	// 此处考虑关闭超时是否有必要
	err := timeoutExecFunc(m.killWaitTTL, func() error {
		// 前置服务关闭
		for _, service := range utils.SliceReverse(m.services) {
			m.stopService(service)
		}

		return nil
	})
	if err != nil {
		m.logger.WARNING("Manager %v stop services error: [%v]", m.name, err.Error())
	}
	m.callback(m.afterServiceShutdown)

	m.logger.INFO("Manager %v stop over", m.name)
}

func (m *Manager) callback(fn Option) {
	if fn != nil {
		fn(m)
	}
}

func (m *Manager) checkSignal() {
	signalChan := make(chan os.Signal, 1)
	if len(m.sigs) == 0 {
		m.sigs = []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM}
	}

	signal.Notify(signalChan, m.sigs...)
	select {
	case exitDetail := <-m.exitChan:
		m.logger.WARNING("Manager %v stop begin, manager receive exitChan msg: [%v]", m.name, exitDetail)
	case sig := <-signalChan:
		m.logger.WARNING("Manager %v stop begin, manager receive closing signal: [%v]", m.name, sig)
	}
}

func timeoutExecFunc(timeout time.Duration, fn func() error) (err error) {
	if timeout <= 0 {
		return fn()
	}
	wait := make(chan error)

	go func() {
		wait <- fn()
	}()

	select {
	case <-time.NewTicker(timeout).C:
		err = errors.New("ExecFunc timeout")
	case err = <-wait:
	}

	return err
}

func (m *Manager) startService(service Service) {
	m.logger.INFO("Service [%v] Init begin", service.GetName())
	if err := service.Init(m); err != nil {
		msg := fmt.Sprintf("Service [%v] Init err: %v", service.GetName(), err)
		m.logger.WARNING(msg)
		panic(msg)
	}

	m.logger.INFO("Service [%v] Init over, OnStart begin", service.GetName())

	if err := service.OnStart(); err != nil {
		msg := fmt.Sprintf("Service [%v] OnStart err: %v", service.GetName(), err)
		m.logger.WARNING(msg)
		panic(msg)
	}
	m.logger.INFO("Service [%v] OnStart over", service.GetName())
}

func (m *Manager) stopService(service Service) {
	m.logger.INFO("Service [%v] OnStop begin", service.GetName())
	if err := service.OnStop(); err != nil {
		m.logger.WARNING("Service [%v] OnStart err: %v", service.GetName(), err)
	}

	m.logger.INFO("Service [%v] OnStop over", service.GetName())
}

func (m *Manager) Stop(msg interface{}) {
	go func() {
		m.exitChan <- msg
	}()
}
