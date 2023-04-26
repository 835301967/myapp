package two_pointer_slice

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"testing"
)

func TestHuiWenSubString(t *testing.T) {
	input := "abnbjjkljljjklfjewfejwklajfgklewgjkewljagklwejgklewjglwwwwwwwwwwwwwwwwwwwwwwww"
	str := getSubLongPendStr(input)
	fmt.Println(str)
	input2 := "abbbbcccccabcdedcbaisgetjkljljkjl"
	str2 := getSubLongPendStr(input2)
	fmt.Println(str2)
}

func pendSlice(str string, i, j int) string {
	for i >= 0 && j < len(str) && str[i] == str[j] {
		i--
		j++
	}
	return str[i+1 : j]
}

func getSubLongPendStr(str string) string {
	temp := ""
	for i := 0; i < len(str)-1; i++ {
		slice := pendSlice(str, i, i)
		slice1 := pendSlice(str, i, i+1)
		if len(temp) < len(slice) {
			temp = slice
		}
		if len(temp) < len(slice1) {
			temp = slice1
		}
	}
	return temp
}

type Order struct {
	Url string
}

func TestUrlJson(t *testing.T) {
	order := Order{Url: "https://gateway-test.myscrm.cn/trade-sak/app/get-signature-url-by-code-and-path?orgcode=fangzhiadmin_test&code=trade-ydjy&path=25/trade-ydjy/fangzhiadmin_test/ydjy/eaa4ff47-4570-471d-8884-2914453f7730"}
	//toString, _ := json.MarshalToString(order)
	consume, _ := gabs.Consume(order)
	fmt.Println(string(consume.EncodeJSON()))
}
