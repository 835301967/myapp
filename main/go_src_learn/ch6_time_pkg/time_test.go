package ch6_time_pkg

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	after := time.Now().After(time.Now().Add(-time.Hour))
	fmt.Println(after)
	date := time.Date(2019, time.February, 28, 10, 10, 10, 0, time.UTC)
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println(format)
	fmt.Println(date)
	parse, _ := time.Parse("2006-01-02 15:04:05", "2019-11-12 13:01:11")
	hourStr := time.Hour.String()
	duration, _ := time.ParseDuration(hourStr)
	parseDuration, _ := time.ParseDuration("10m0s")
	fmt.Println(parse)
	fmt.Println(duration)
	fmt.Println(hourStr)
	fmt.Println(parseDuration.Seconds())

}

// 关于Timer和Ticker用法参考
//https://juejin.cn/post/6884914839308533774
func TestTimeTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 1)
	fmt.Println("start a ticker")
	now := time.Now()
	go func() {
		for {
			fmt.Println("do a ticker")
			since := time.Since(now)
			fmt.Println(since)
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("stop ticker")
	ticker.Stop()
	time.Sleep(time.Second * 5)
	since1 := time.Since(now)
	fmt.Println(since1)
}

func TestTimeTick(t *testing.T) {
	timer := time.NewTimer(time.Second * 1)
	now := time.Now()
	for i := 0; i < 10; i++ {
		<-timer.C
		fmt.Println("do a timer")
		fmt.Println(time.Since(now))
		timer.Reset(time.Second * 2)
	}
	fmt.Println("done")
}
