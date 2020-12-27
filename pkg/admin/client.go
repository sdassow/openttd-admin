package admin

import (
	"net"
)

type Client struct {
	conn net.Conn
}

func Connect(addr, password, name, version string) (*Client, error) {
	c := &Client{}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	tcpconn, ok := conn.(*net.TCPConn)
	if ok {
		tcpconn.SetNoDelay(true)
	}

	c.conn = conn

	c.Send(NewAdminJoin(password, name, version))

	return c, nil
}

func (c *Client) Send(p *Packet) (int, error) {
	return c.conn.Write(p.Marshal())
}

func (c *Client) ReadMessage() (interface{}, error) {
	p, err := ReadPacket(c.conn)
	if err != nil {
		return nil, err
	}
	return ParseMessage(p)
}

func (c *Client) Close() error {
	c.Send(NewAdminQuit())
	return c.conn.Close()
}
