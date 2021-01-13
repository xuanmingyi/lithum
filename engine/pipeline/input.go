package pipeline

import "engine/models"

type Input interface {
	Start(chan models.Message)
}
