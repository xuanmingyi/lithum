package config

type Attribute struct {
	Name    string `yaml:"name"`
	Display string `yaml:"display"`
	Tag     string `yaml:"tag"`
	Values  map[string]interface{}
}

type Field struct {
	Name string `yaml:"name"`
}

type Dialog struct {
	Title  string  `yaml:"title"`
	Fields []Field `yaml:"fields"`
}

type Action struct {
	Name    string `yaml:"name"`
	Display string `yaml:"display"`
	Class   string `yaml:"class"`
	Type    string `yaml:"type"`
	Dialog  Dialog `yaml:"dialog"`
}

type Column struct {
	Name    string `yaml:"name"`
	Display string `yaml:"display"`
	Align   string `yaml:"align"`
	Width   int    `yaml:"width"`
}

type Table struct {
	Columns      []Column `yaml:"columns"`
	TableActions []Action `yaml:"actions"`
	RowActions   []Action `yaml:"row_actions"`
}

type Model struct {
	Name       string      `yaml:"name"`
	Title      string      `yaml:"title"`
	Attributes []Attribute `yaml:"attributes"`
	Table      Table       `yaml:"table"`
}

func (model *Model) GetAction(action_name string) (action *Action) {
	for _, act := range model.Table.TableActions {
		if act.Name == action_name {
			return &act
		}
	}

	for _, act := range model.Table.RowActions {
		if act.Name == action_name {
			return &act
		}
	}

	return nil
}

func LoadModel(name string) (m *Model, err error) {
	m = &Model{Name: name}
	return m, nil
}
