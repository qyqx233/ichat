package api

import (
	"context"
)

type ClientCfg struct {
	Name, Ak, Sk string
}

// 泛型接口
type ClientProcessor[R any, S any] interface {
	Process(ctx context.Context, r R) (s S, err error)
}

type A int

func Process[R, S any](ctx context.Context, p []ClientProcessor[R, S], rq R) (S, error) {
	c := p[0]
	return c.Process(ctx, rq)
}

// func (c []ClientProcessor[R, S]) Process(r R) (s T) {
// 	return nil, nil
// }

func NewAll(cfgs []ClientCfg) {
	for _, c := range cfgs {
		if c.Name == "openai" {
			OpenAIClient = append(OpenAIClient, newOpenAI(c))
		}
	}
}

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}
