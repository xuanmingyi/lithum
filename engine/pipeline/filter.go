package pipeline

type Filter interface {
	Start(chan Message, chan Message)
}
