package task

type Task interface {
	Run()

	GetName() string
}
