package main

import (
	"fmt"
	"sync"
)

var locker = sync.Mutex{}
var count = 0

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go func() {
		for i := 0; i < 1000000; i++ {
			countAdd()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			countAdd()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("count：%v\n", count)
}

func countAdd() {
	lockExec(func() any {
		count++
		return count
	})
}

func lockExec(f func() any) any {
	locker.Lock()
	defer locker.Unlock()
	return f()
}
