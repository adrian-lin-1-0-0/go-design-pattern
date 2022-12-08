package abstractfactory

import (
	"testing"
)

func TestGetSportsFactory(t *testing.T) {
	type args struct {
		brand string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Test ISportsFactory : adidas", args: args{brand: "adidas"}, want: "adidas"},
		{name: "Test ISportsFactory : nike", args: args{brand: "nike"}, want: "nike"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factory, err := GetSportsFactory(tt.args.brand)
			shoeLogo := factory.MakeShoe().GetLogo()
			shirtLogo := factory.MakeShirt().GetLogo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSportsFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if shoeLogo != tt.want || shirtLogo != tt.want {
				t.Errorf("GetSportsFactory() = %v,%v, want %v", shoeLogo, shirtLogo, tt.want)
			}
		})
	}
}
