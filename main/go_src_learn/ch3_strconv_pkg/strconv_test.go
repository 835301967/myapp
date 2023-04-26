package ch3_strconv_pkg

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestStrInt(t *testing.T) {
	bytes := make([]byte, 0)
	bytes = strconv.AppendInt(bytes, 10, 10)
	bytes = strconv.AppendInt(bytes, 3242, 10)
	bytes = strconv.AppendInt(bytes, 324, 10)
	fmt.Println(string(bytes))
	bytes = strconv.AppendQuote(bytes, "jkglewjgklea")
	fmt.Println(string(bytes))
	formatInt := strconv.FormatInt(3423, 10)
	fmt.Println(formatInt)

}

func TestStrParse(t *testing.T) {
	float, _ := strconv.ParseFloat("234324.324324", 64)
	fmt.Println(float)
	fmt.Println(reflect.TypeOf(float))
	quote := strconv.Quote("jkljfewfew")
	fmt.Println(quote)
	i, _ := strconv.ParseInt("432432", 10, 64)
	fmt.Println(i)
	var input = 123213.3213
	intS := int64(input)
	fmt.Println(intS)
	var input1 = 123213.5213
	intS2 := int64(input1)
	fmt.Println(intS2)
}
