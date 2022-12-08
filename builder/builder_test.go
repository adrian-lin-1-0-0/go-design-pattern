package builder

import (
	"testing"
)

func TestGetBuilder(t *testing.T) {
	type args struct {
		builderTpye string
	}
	tests := []struct {
		name       string
		args       args
		wantDoor   string
		wantWindow string
	}{
		{name: "test redBuilder", args: args{builderTpye: "red"}, wantDoor: "red", wantWindow: "red"},
		{name: "test blueBuilder", args: args{builderTpye: "blue"}, wantDoor: "blue", wantWindow: "blue"},
	}

	director := NewDirector(nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := GetBuilder(tt.args.builderTpye)
			director.SetBuilder(builder)
			house := director.buildHouse()
			if house.GetDoorColor() != tt.wantDoor {
				t.Errorf("buildHouse().door = %v, want %v", house.GetDoorColor(), tt.wantDoor)
			}

			if house.GetWindowColor() != tt.wantWindow {
				t.Errorf("buildHouse().window = %v, want %v", house.GetWindowColor(), tt.wantWindow)
			}
		})
	}
}
