package ch4_sync_pkg

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncLock(t *testing.T) {
	mutex := &sync.Mutex{}
	mutex.Lock()
	mutex.Unlock()
}

func TestSyncWaitGroup(t *testing.T) {
	waitGroup := sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			fmt.Println("this is  a grountine，do some thing wait a second")
			time.Sleep(time.Second)
			fmt.Printf("this is  a grountine:%d \n", i)
		}(i)
	}
	waitGroup.Wait()
	fmt.Println("program is done")
}

func TestSyncCond(t *testing.T) {
	cond := &sync.Cond{}
	cond = sync.NewCond(new(sync.Mutex))
	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Wait()
	fmt.Println("cond wait")
	go func() {
		cond.Signal()
		fmt.Println("cond signal")
	}()
	cond.Broadcast()
}

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{}
	pool.Put(Some{
		Name:          "123",
		Age:           123,
		LastLoginTime: time.Now(),
	})
	pool.Put(Some{
		Name:          "4324",
		Age:           11,
		LastLoginTime: time.Now(),
	})
	pool.Put(Some{
		Name:          "hgreg",
		Age:           34,
		LastLoginTime: time.Now(),
	})
	get := pool.Get()
	get2 := pool.Get()
	get3 := pool.Get()
	get4 := pool.Get()
	fmt.Println(get)
	fmt.Println(get2)
	fmt.Println(get3)
	fmt.Println(get4)
}

func TestSyncRWLock(t *testing.T) {
	s := sync.Map{}
	group := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func(i int) {
			defer group.Done()
			s.Store(i, "jkflejwfel")
		}(i)
	}
	group.Wait()
	count := 0
	s.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Println(count)
}

func TestSyncSemaphore(t *testing.T) {

}

func TestSyncOnce(t *testing.T) {
	once := sync.Once{}
	for i := 0; i < 100; i++ {
		once.Do(func() {
			fmt.Println("jkljklj")
		})
	}
	time.Sleep(time.Second * 2)
}

type Some struct {
	Name          string
	Age           int
	LastLoginTime time.Time
}
