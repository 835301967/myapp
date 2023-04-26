package main

import (
	"fmt"
	"testing"
)



func TestChant(t *testing.T) {
	stringChan := make(chan int)
	go func() {
		for i:=10;i<20;i++{
			stringChan<- i
		}
	}()
	for  j:=range stringChan{
		fmt.Printf("输出数据：%d \n",j)
	}
	fmt.Println("done")
	return
}

func TestSlice(t *testing.T)  {
	slice := make([]string, 5)
	slice[0] = "1"
	slice[1] = "2134"
	slice[2] = "sfwef"
	slice[3] = "sfwef"
	slice[4] = "sfwef"

	sliceLen := len(slice)
	fmt.Printf("len:%d\n",sliceLen)
	fmt.Printf("slice:%+v\n",slice)

	slice2 := make([]string, 10)
	slice2[0] = "2134"


	sliceLen2 := len(slice)
	fmt.Printf("len2:%d\n",sliceLen2)
	fmt.Printf("slice2:%+v\n",slice2)

	slice3 := slice[1:3]


	sliceLen3 := len(slice3)
	fmt.Printf("len3:%d\n",sliceLen3)
	fmt.Printf("slice3:%+v\n",slice3)

}
