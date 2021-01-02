package model

type Option struct {
	Label string `yaml:"label" json:"label"`
	Value string `yaml:"value" json:"value"`
}

type Attribute struct {
	Name    string   `yaml:"name" json:"name"`
	Display string   `yaml:"display" json:"display"`
	Tag     string   `yaml:"tag" json:"tag"`
	Options []Option `yaml:"options" json:"options"`
}

type Model struct {
	Name      string      `yaml:"name" json:"name"`
	Attribute []Attribute `yaml:"attributes" json:"attributes"`
}

type Column struct {
	Name string `yaml:"name" json:"name"`
}

type Field struct {
	Name string `yaml:"name" json:"name"`
}

type Dialog struct {
	Title  string  `yaml:"title" json:"title"`
	Fields []Field `yaml:"fields" json:"fields"`
}

type Action struct {
	Name    string `yaml:"name" json:"name"`
	Display string `yaml:"display" json:"display"`
	Type    string `yaml:"type" json:"type"`
	Dialog  Dialog `yaml:"dialog" json:"dialog"`
}

type Table struct {
	Columns    []Column `yaml:"columns" json:"columns"`
	Actions    []Action `yaml:"actions" json:"actions"`
	RowActions []Action `yaml:"row_actions" json:"row_actions"`
}

type MetaData struct {
	Model Model `yaml:"model" json:"model"`
	Table Table `json:"table" json:"table"`
}