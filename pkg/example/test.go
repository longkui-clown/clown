package example

import (
	"fmt"
	"time"

	"github.com/longkui-clown/clown/core"
)

type TestApp struct {
	core.BaseService
	_m *core.Manager
}

func NewTestApp(name string) core.Service {
	testApp := &TestApp{}
	testApp.BaseService.SetName(name)
	return testApp
}

func (ta *TestApp) Init(m *core.Manager) error {
	ta._m = m
	return nil
}

func (ta *TestApp) OnStart() error {
	go func() {
		time.Sleep(time.Second * 5)
		ta._m.Stop(fmt.Sprintf("%v manuam stop manager ...", ta.GetName()))
	}()

	return nil
}

func (ta *TestApp) OnStop() error {
	ta._m.Logger().INFO("@@@@ %v OnStop calback", ta.GetName())
	return nil
}
