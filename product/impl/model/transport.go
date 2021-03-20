package model

type Transport struct {
	Name string
}

func (g *Transport) SetName(name string) {
	g.Name = name
}

func (g *Transport) GetName() string {
	return g.Name
}