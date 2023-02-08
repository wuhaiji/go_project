package test_common

import (
	"fmt"
	"testing"
)

func TestGOID(t *testing.T) {
	id := curGoroutineID()
	fmt.Printf("go id = %v\n", id)
}
