package pipeline

import "engine/models"

type Output interface {
	Start(chan models.Message)
}
