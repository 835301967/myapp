package main

import (
	"fmt"
	"sync"
)

func main() {
	ms := make([]*M, 10)
	ms[0] = &M{
		X: "Self",
		N: 123,
	}
	ms[1] = &M{
		X: "fewg",
		N: 166,
	}
	ms[2] = &M{
		X: "ge",
		N: 66,
	}
	waitGroup := sync.WaitGroup{}
	for _,m:= range ms{
		fmt.Printf("对象值：%+v \n",m)
		waitGroup.Add(1)
		go func(m *M) {
			//fmt.Printf("对象值：%+v \n",m)
			waitGroup.Done()
		}(m)
	}
	waitGroup.Wait()
}

type M struct {
	N int
	X string
}