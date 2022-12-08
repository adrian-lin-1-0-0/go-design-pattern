package adapter

type Clinet struct {
}

func (c *Clinet) Connect2Server(con Connection) string {
	return con.Tcp()
}

type Connection interface {
	Tcp() string
}

type TcpConnection struct {
}

func (tcp *TcpConnection) Tcp() string {
	return "tcp"
}

type UdpConnection struct {
}

func (udp *UdpConnection) Udp() string {
	return "udp"
}

type UdpAdapter struct {
	UdpConnection *UdpConnection
}

func (ua *UdpAdapter) Tcp() string {
	return ua.UdpConnection.Udp()
}
