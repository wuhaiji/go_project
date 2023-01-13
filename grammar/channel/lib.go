package channel

import "fmt"

// Channel 对系统channel的包装,使得更加方便使用
type Channel[T any] struct {
	channel chan T
}

func (c Channel[T]) String() string {
	sprintf := fmt.Sprintf("%v", c.channel)
	return sprintf
}

func NewChannel[T any](size int) *Channel[T] {
	if size < 0 {
		size = 0
	}
	ch := make(chan T, size)
	return &Channel[T]{channel: ch}
}

func NewNoBufChannel[T any]() *Channel[T] {
	ch := make(chan T, 0)
	return &Channel[T]{channel: ch}
}

func (c Channel[T]) Send(data T) {
	c.channel <- data
}

func (c Channel[T]) Recv() T {
	data := <-c.channel
	return data
}

func (c Channel[T]) Close() {
	close(c.channel)
}
