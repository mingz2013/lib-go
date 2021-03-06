package tcp

import (
	"bufio"
	"net"
)

type Conn struct {
	server *Server

	removeAddr string

	r    *connReader
	bufr *bufio.Reader
	bufw *bufio.Writer

	rwc net.Conn

	extra int // 一个额外的指针类型的字段字段，可用于设置一些数据，做反向绑定
}

func (c Conn) Serve() error {

	for {
		buffer, err := c.readBuffer()

		if err != nil {
			return err
		}

		c.server.Handler.Serve(c, buffer)

	}

}

func (c *Conn) readBuffer() (b []byte, e error) {
	//n,e:=c.rwc.Read(b)
	//if
	c.bufr.Read(b)

	return
}

func (c Conn) WriteBuffer(buf []byte) {

}

//func (c *Conn) readRequest() (net2.Request, error) {
//	return net2.Request{}, nil
//}
//
//func (c Conn) WriteResponse(resp net2.Response) {
//
//}

func (c Conn) GetExtra() int {
	return c.extra
}

func (c Conn) SetExtra(e int) {
	c.extra = e
}
