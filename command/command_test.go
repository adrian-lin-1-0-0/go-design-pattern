package command

import "testing"

func TestPlugButton(t *testing.T) {
	plug := &Plug{}
	onCommand := &OnCommand{device: plug}
	onButton := &Button{command: onCommand}
	offCommand := &OffCommand{device: plug}
	offButton := &Button{command: offCommand}

	type fields struct {
		command func()
		status  func() bool
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Test Plug On", fields: fields{
			command: onButton.Press,
			status:  plug.IsRunning}, want: true},
		{name: "Test Plug Off", fields: fields{
			command: offButton.Press,
			status:  plug.IsRunning}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.command()
			got := tt.fields.status()
			if got != tt.want {
				t.Errorf("Button.Press() got = %v, want %v", got, tt.want)

			}
		})
	}

}
