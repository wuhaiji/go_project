package channel

import (
	"fmt"
	"testing"
)

func TestChannelSend(t *testing.T) {
	var channel = NewChannel[int](0)
	fmt.Printf("channel=%v", channel.String())
}
