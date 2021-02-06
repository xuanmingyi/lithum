package config

type Option struct {
	Display string `yaml:"display"`
	Value   string `yaml:"value"`
}

type Attribute struct {
	Name    string   `yaml:"name"`
	Display string   `yaml:"display"`
	Tag     string   `yaml:"tag"`
	Options []Option `yaml:"options"`
}

type Field struct {
	Name     string     `yaml:"name"`
	Disabled bool       `yaml:"disabled"`
	Attr     *Attribute `yaml:"-"`
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
