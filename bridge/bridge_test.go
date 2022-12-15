package bridge

import "testing"

func TestMac_Print(t *testing.T) {
	type fields struct {
		printer  Printer
		computer Computer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "test mac hp", fields: fields{printer: &Hp{}, computer: &Mac{}}, want: "Mac HP"},
		{name: "test mac epson", fields: fields{printer: &Epson{}, computer: &Mac{}}, want: "Mac EPSON"},
		{name: "test linux hp", fields: fields{printer: &Hp{}, computer: &Linux{}}, want: "Linux HP"},
		{name: "test windows epson", fields: fields{printer: &Epson{}, computer: &Linux{}}, want: "Linux EPSON"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.computer.SetPrinter(tt.fields.printer)
			got := tt.fields.computer.Print()
			if got != tt.want {
				t.Errorf("Computer.Print() got = %v, want %v", got, tt.want)
			}

		})
	}
}
