package proxy

import "testing"

func TestProxy_HandleRequest(t *testing.T) {

	type args struct {
		url    string
		method string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 1", args: args{url: "/app/status", method: "GET"}, want: 200},
		{name: "Test 2", args: args{url: "/app/status", method: "GET"}, want: 403},
	}
	p := NewProxyServer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := p.HandleRequest(tt.args.url, tt.args.method)
			if got != tt.want {
				t.Errorf("Proxy.HandleRequest() got = %v, want %v", got, tt.want)
			}

		})
	}
}
