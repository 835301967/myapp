package dom

import (
	"fmt"
	"myapp/main/common"
	"reflect"
	"testing"
)

func TestConn(t *testing.T) {
	c := new(common.Conn)
	jk:=common.Conn{
		Name: "jkljkl",
	}
	c.SetName("fwjekl")
	fmt.Printf("struct:%+v \n",c)
	c.SetName2("pointer")
	fmt.Printf("pointer:%+v \n",c)
	p:=*c
	p.SetName2("jklj")
	fmt.Printf("test:%+v \n",p)
	fmt.Printf("p type:%+v \n",reflect.TypeOf(p))
	fmt.Printf("c type:%+v \n",reflect.TypeOf(c))
	fmt.Printf("jk type:%+v \n",reflect.TypeOf(jk))
	fmt.Printf("struce:%+v \n",jk)
}
