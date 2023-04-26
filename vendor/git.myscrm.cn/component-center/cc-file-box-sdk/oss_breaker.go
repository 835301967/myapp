package file_box

import (
	"context"
)
import "github.com/afex/hystrix-go/hystrix"

type breaker interface {
	Call(context.Context, func(ctx context.Context) error) error
}

type hystrixBreaker struct {
	name string
}

func newHystrixBreaker(name string) *hystrixBreaker {
	hystrix.ConfigureCommand(name, hystrix.CommandConfig{
		Timeout:                1000 * 5, // 5s
		MaxConcurrentRequests:  2000,
		RequestVolumeThreshold: 10,
		SleepWindow:            1000 * 10, // 10s
		ErrorPercentThreshold:  50,
	})
	return &hystrixBreaker{name: name}
}

func (m *hystrixBreaker) Call(ctx context.Context, f func(ctx context.Context) error) error {
	return hystrix.Do(m.name, func() error {
		return f(ctx)
	}, func(err error) error {
		return err
	})
}

func newOssManagerBreaker() breaker {
	return newHystrixBreaker("oss-manager")
}
