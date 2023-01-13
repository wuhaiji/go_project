package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func ExtractGID(s []byte) int64 {
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

// Parse the goid from runtime.Stack() output. Slow, but it works.
func getSlow() int64 {
	var buf [64]byte
	stack := runtime.Stack(buf[:], false)
	s := buf[:stack]
	return ExtractGID(s)
}

func GetGID() int64 {
	return getSlow()
}

func main() {
	slow := getSlow()
	var wg = sync.WaitGroup{}
	wg.Add(1000)
	fmt.Printf("%v", slow)
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Printf("%v\n", getSlow())
			wg.Done()
		}()
	}
	wg.Wait()
}
