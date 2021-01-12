package pipeline

type Output interface {
	Start(chan Message)
}
