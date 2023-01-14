package channel

import (
	"errors"
	"fmt"
	"reflect"
)

// channel 对系统channel的包装,使得更加方便使用
type channel struct {
	channel chan any
	state   byte
	Ty      reflect.Type
}
type Channel = *channel

var (
	OPENED = 0
	CLOSED = 1
)

func NewBufChannel(size int) Channel {
	if size < 0 {
		size = 0
	}
	ch := make(chan any, size)
	return &channel{channel: ch}
}

func NewChannel() Channel {
	return NewBufChannel(0)
}

const errorDataTypeInconsistent = "datatype is not inconsistent ,channel type:%v,data type:%v"

func (c *channel) Send(data any) error {
	dataType := reflect.TypeOf(data)
	if c.Ty == nil {
		c.Ty = dataType
	} else if c.Ty != dataType {
		return fmt.Errorf(errorDataTypeInconsistent, c.Ty, dataType)
	}
	c.channel <- data
	return nil
}

func (c *channel) Recv() (data any, err error) {
	data, ok := <-c.channel

	if ok {
		dataType := reflect.TypeOf(data)
		if c.Ty != dataType {
			return nil, fmt.Errorf(errorDataTypeInconsistent, c.Ty, dataType)
		}
		return data, nil
	} else {
		return data, errors.New("channel is closed")
	}
}

func (c *channel) Range(f func(any)) {
	for data := range c.channel {
		f(data)
	}
}

func (c *channel) Close() {

	close(c.channel)
}

func (c *channel) Len() int {
	return len(c.channel)
}

func (c *channel) Cap() int {
	return cap(c.channel)
}

func (c *channel) Inner() *chan any {
	return &c.channel
}
func (c *channel) String() string {
	return fmt.Sprintf("%v", c.channel)
}
