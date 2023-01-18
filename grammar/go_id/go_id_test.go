package go_id

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGoID(t *testing.T) {
	go func() {
		id := GetGoroutineID()
		fmt.Println("goroutine id:" + strconv.FormatInt(id, 10))
	}()
}
