package config

type Attribute struct {
	Name    string `yaml:"name"`
	Display string `yaml:"display"`
	Tag     string `yaml:"tag"`
	Values  map[string]interface{}
}

type Action struct {
	Name    string `yaml:"name"`
	Display string `yaml:"display"`
	Type    string `yaml:"type"`
}

type Table struct {
	TableActions []Action `yaml:"actions"`
	RowActions   []Action `yaml:"row_actions"`
}

type Model struct {
	Name       string      `yaml:"name"`
	Title      string      `yaml:"title"`
	Attributes []Attribute `yaml:"attributes"`
}

func LoadModel(name string) (m *Model, err error) {
	m = &Model{Name: name}
	return m, nil
}
