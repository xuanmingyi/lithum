package config

import (
	"database/sql"
	"fmt"
	"strings"
)

type Field struct {
	Display string
	TypeName string
}

func LoadField(comment string) (field *Field) {
	if len(comment) == 0 {
		return nil
	}

	field = new(Field)
	for _, attr := range strings.Split(comment, ";") {
		kvs := strings.Split(attr, ":")
		key := kvs[0]
		value := kvs[1]
		switch key {
		case "display":
			field.Display = value
		case "type":
			field.TypeName = key
		}
	}
	return field
}

type Model struct {
	Name string
	Fields map[string]*Field
}

func (m *Model) Load() (err error){
	var rows *sql.Rows
	if Config.DB == nil {
		panic("error db")
	}

	rows, err = Config.DB.Query("SELECT COLUMN_NAME, COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA=? and TABLE_NAME=?", "sys", m.Name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var name string
	var comment string

	if m.Fields == nil {
		m.Fields = make(map[string]*Field)
	}

	for rows.Next() {
		err = rows.Scan(&name, &comment)
		if err != nil {
			panic(err)
		}
		m.Fields[name] = LoadField(comment)
		fmt.Println(name, comment)
	}
	return nil
}

func LoadModel(name string) (m *Model, err error) {
	m = &Model{Name: name}

	err = m.Load()
	if err != nil {
		panic(err)
	}
	return m, nil
}


