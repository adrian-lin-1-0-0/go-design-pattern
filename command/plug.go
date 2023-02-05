package command

type Plug struct {
	isRunning bool
}

func (plug *Plug) IsRunning() bool {
	return plug.isRunning
}

func (plug *Plug) On() {
	plug.isRunning = true
}

func (plug *Plug) Off() {
	plug.isRunning = false
}
