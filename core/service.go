package core

type Service interface {
	Init(*Manager) error
	GetName() string
	SetName(name string) Service
	OnStart() error
	OnStop() error
}

var _ Service = (*BaseService)(nil)

type BaseService struct {
	name string
}

func (*BaseService) Init(*Manager) error {
	return nil
}

func (base *BaseService) GetName() string {
	return base.name
}

func (base *BaseService) SetName(name string) Service {
	base.name = name
	return base
}

func (*BaseService) OnStart() error {
	return nil
}

func (*BaseService) OnStop() error {
	return nil
}
