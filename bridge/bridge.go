package bridge

type Computer interface {
	Print() string
	SetPrinter(Printer)
}

type Printer interface {
	PrintFile(string) string
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() string {
	return m.printer.PrintFile("Mac")
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Linux struct {
	printer Printer
}

func (l *Linux) Print() string {
	return l.printer.PrintFile("Linux")
}

func (l *Linux) SetPrinter(p Printer) {
	l.printer = p
}

type Epson struct {
}

func (p *Epson) PrintFile(s string) string {
	return s + " EPSON"
}

type Hp struct {
}

func (p *Hp) PrintFile(s string) string {
	return s + " HP"
}
