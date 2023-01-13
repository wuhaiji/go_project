package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var rt = NewReentrantLock()
var count = 0

var wg = initWg()
var gCount = 3

func initWg() *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(gCount)
	return wg
}
func main() {
	go func() {
		for true {
			time.Sleep(time.Second)
			fmt.Printf("哈哈哈\n")
		}
		wg.Done()
	}()
	go add()
	go func() {
		rt.Lock()
		defer rt.Unlock()
		add()
	}()
	wg.Wait()
	fmt.Printf("count:%v", count)

}

func add() {
	for i := 0; i < 1000; i++ {
		addOne()
	}
	wg.Done()
}

func addOne() {
	rt.Lock()
	defer rt.Unlock()
	count++
}

type ReentrantLock struct {
	lock      *sync.Mutex
	cond      *sync.Cond
	recursion int32
	host      int64
}

func NewReentrantLock() *ReentrantLock {
	res := &ReentrantLock{
		lock:      new(sync.Mutex),
		recursion: 0,
		host:      0,
	}
	res.cond = sync.NewCond(res.lock)
	return res
}

func (rt *ReentrantLock) Lock() {
	rt.lock.Lock()
	defer rt.lock.Unlock()
	fmt.Printf("rt:%v\n", rt)

	id := GetGoroutineID()
	if rt.host == id {
		rt.recursion++
		return
	}

	for rt.recursion != 0 {
		rt.cond.Wait()
	}
	rt.host = id
	rt.recursion = 1

}

func (rt *ReentrantLock) Unlock() {
	rt.lock.Lock()
	defer rt.lock.Unlock()

	if rt.recursion == 0 || rt.host != GetGoroutineID() {
		err := fmt.Sprintf(
			"the wrong call host: (%d); current_id: %d; recursion: %d",
			rt.host, GetGoroutineID(),
			rt.recursion,
		)
		panic(err)
	}

	rt.recursion--
	if rt.recursion == 0 {
		rt.cond.Signal()
	}
}

func GetGoroutineID() int64 {
	return getSlow()
}

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
