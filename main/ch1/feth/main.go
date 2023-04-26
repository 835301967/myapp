package main

import (
	"fmt"
	"os"
	_ "strings"
	"time"
)

func main()  {
	fmt.Printf("%+v",os.Args)
	mymap:=make(map[string]string,0)
	mymap["jklj"] = "rewr"
	go func() {
		fmt.Printf("jklj \n")
		fmt.Println("----------------------------------")
		defer func() {
			recover()
		}()
		go func() {
			defer func() {
				if recover() != nil {
					fmt.Printf("捕捉到了一个")
				}
			}()

			panic("this is a panic")
		}()

	}()
	fmt.Println(mymap)
	time.Sleep(time.Second*5)
}

