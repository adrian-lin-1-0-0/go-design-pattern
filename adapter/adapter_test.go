package adapter

import "testing"

func TestClinet_Connect2Server(t *testing.T) {
	client := &Clinet{}
	tcp := &TcpConnection{}
	got := client.Connect2Server(tcp)
	want := "tcp"
	if got != want {
		t.Errorf("Connect2Server() = %v, want %v", got, want)
	}
	udp := &UdpConnection{}
	udpAdapter := &UdpAdapter{
		UdpConnection: udp,
	}

	got = client.Connect2Server(udpAdapter)
	want = "udp"
	if got != want {
		t.Errorf("Connect2Server() = %v, want %v", got, want)
	}

}
