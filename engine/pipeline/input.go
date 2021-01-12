package pipeline

type Input interface {
	Start(chan Message)
}
