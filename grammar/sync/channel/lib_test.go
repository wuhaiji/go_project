package channel

import (
	"fmt"
	"testing"
)

func TestChannelSend(t *testing.T) {
	var channel = NewChannel[int]()
	fmt.Printf("channel=%v", channel.String())
}
