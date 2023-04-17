package components

type Component interface {
	Init()
}

type Module struct {
	component []Component
}

func New() *Module {
	return &Module{}
}

func (m *Module) Add(component ...Component) {
	m.component = append(m.component, component...)
}
