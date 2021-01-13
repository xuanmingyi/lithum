package pipeline

import (
	"engine/models"
	"fmt"
)

type Filter struct {
}

func (f *Filter) Start(i chan models.Message, o chan models.Message) {
	var m models.Message

	for {
		fmt.Println("filter init")
		m = <-i
		fmt.Println("filter init222")
		fmt.Println(m.Body)
		o <- m
	}
}
