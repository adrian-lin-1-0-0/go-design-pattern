package command

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.On()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.Off()
}
