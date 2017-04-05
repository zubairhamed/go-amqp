package performatives

import "net"

type Performative interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}


func SendPerformative(c net.Conn, p Performative) (int, error) {
	b, err := p.Encode()
	if err != nil {
		panic(err.Error())
	}

	c.Write(b)

	return 0, nil
}
