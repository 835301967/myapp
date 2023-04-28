package ch5_atomic_pkg

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var input int64 = 0
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			addInt64 := atomic.AddInt64(&input, 1)
			fmt.Println(addInt64)
		}()
	}
	waitGroup.Wait()
	fmt.Println(input)
}

func TestAtomicAddFloat(t *testing.T) {
	var input int64 = 0
	swapped := atomic.CompareAndSwapInt64(&input, 23, 23)
	fmt.Println(swapped)
	fmt.Println(input)
	var input2 int32 = 10
	atomic.StoreInt32(&input2, 1000)
	fmt.Println(input2)
}

func TestAtomicAdd(t *testing.T) {
	var in int32 = 234
	val := atomic.LoadInt32(&in)
	fmt.Println(val)
}
