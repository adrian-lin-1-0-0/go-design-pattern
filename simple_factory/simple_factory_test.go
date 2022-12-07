package simplefactory

import "testing"

func TestAnimal(t *testing.T) {
	type args struct {
		animalName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test human", args: args{animalName: ""}, want: "hi"},
		{name: "test cat", args: args{animalName: "cat"}, want: "meow"},
		{name: "test dog", args: args{animalName: "dog"}, want: "woof"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if animal := NewAnimal(tt.args.animalName); animal.Say() != tt.want {
				t.Errorf("animal.Say() = %v, want %v", animal.Say(), tt.want)
			}
		})
	}
}
