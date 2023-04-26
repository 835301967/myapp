package common

type Conn struct {
	Name string
	Age int
}

func (conn Conn) SetName(name string)  {
	conn.Name = name
}

func (conn *Conn) SetName2(name string) {
	conn.Name = name
}
