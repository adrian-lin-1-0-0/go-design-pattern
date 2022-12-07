package factorymethod

import "testing"

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	plusOperatorFactory := PlusOperatorFactory{}
	if compute(plusOperatorFactory, 1, 2) != 3 {
		t.Fatal()
	}

	type args struct {
		factory OperatorFactory
		a, b    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test PlusOperatorFactory",
			args: args{
				factory: PlusOperatorFactory{},
				a:       1,
				b:       2,
			},
			want: 3,
		},
		{
			name: "test MinusOperatorFactory",
			args: args{
				factory: MinusOperatorFactory{},
				a:       1,
				b:       2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if result := compute(tt.args.factory, tt.args.a, tt.args.b); result != tt.want {
				t.Errorf("Result() = %v, want %v", result, tt.want)
			}
		})
	}
}
