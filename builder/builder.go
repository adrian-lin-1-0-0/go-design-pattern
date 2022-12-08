package builder

type House struct {
	window string
	door   string
}

func (h *House) GetDoorColor() string {
	return h.door
}

func (h *House) GetWindowColor() string {
	return h.window
}

type Builder interface {
	setWindowColor()
	setDoorColor()
	GetHouse() House
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorColor()
	d.builder.setWindowColor()
	return d.builder.GetHouse()
}

type RedBuilder struct {
	window string
	door   string
}

func newRedBuilder() *RedBuilder {
	return &RedBuilder{}
}

func (rb *RedBuilder) setWindowColor() {
	rb.window = "red"
}

func (rb *RedBuilder) setDoorColor() {
	rb.door = "red"
}

func (rb *RedBuilder) GetHouse() House {
	return House{
		door:   rb.door,
		window: rb.window,
	}
}

type BlueBuilder struct {
	window string
	door   string
}

func newBlueBuilder() *BlueBuilder {
	return &BlueBuilder{}
}

func (bb *BlueBuilder) setWindowColor() {
	bb.window = "blue"
}

func (bb *BlueBuilder) setDoorColor() {
	bb.door = "blue"
}

func (bb *BlueBuilder) GetHouse() House {
	return House{
		door:   bb.door,
		window: bb.window,
	}
}

func GetBuilder(builderTpye string) Builder {
	switch builderTpye {
	case "red":
		return newRedBuilder()
	case "blue":
		return newBlueBuilder()
	default:
		return nil
	}
}
