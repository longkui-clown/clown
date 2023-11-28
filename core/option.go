package core

import (
	"os"
	"syscall"
	"time"
)

type Option func(*Manager)

func Name(name string) Option {
	return func(m *Manager) {
		m.name = name
	}
}

func KillWaitTTL(ttl int64) Option {
	return func(m *Manager) {
		m.killWaitTTL = (time.Duration)(ttl) * time.Millisecond
	}
}

func BeforeStartup(fn Option) Option {
	return func(m *Manager) {
		m.beforeStartup = fn
	}
}

func AfterServiceStartup(fn Option) Option {
	return func(m *Manager) {
		m.afterServiceStartup = fn
	}
}

func AfterAppStartup(fn Option) Option {
	return func(m *Manager) {
		m.afterAppStartup = fn
	}
}

func BeforeShutdown(fn Option) Option {
	return func(m *Manager) {
		m.beforeShutdown = fn
	}
}

func AfterServiceShutdown(fn Option) Option {
	return func(m *Manager) {
		m.afterServiceShutdown = fn
	}
}

func AfterAppShutdown(fn Option) Option {
	return func(m *Manager) {
		m.afterAppShutdown = fn
	}
}

func Services(services ...Service) Option {
	return func(m *Manager) {
		m.services = services
	}
}

func App(app Service) Option {
	return func(m *Manager) {
		m.app = app
	}
}

func SetLogger(l Logger) Option {
	return func(m *Manager) {
		m.logger = l
	}
}

func WaitExitSignal(sigs ...os.Signal) Option {
	return func(m *Manager) {
		if len(sigs) == 0 {
			sigs = []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM}
		}

		m.sigs = sigs
	}
}
