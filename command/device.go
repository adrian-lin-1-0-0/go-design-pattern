package command

type Device interface {
	On()
	Off()
	IsRunning() bool
}
