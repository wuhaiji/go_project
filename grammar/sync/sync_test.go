package sync

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	var m = make(map[string]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m[key] = n
			fmt.Printf("k=:%v,v:=%v\n", key, m[key])
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test2(t *testing.T) {
	var m = sync.Map{}
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 存储key-value
			value, _ := m.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test3(t *testing.T) {
	var counter int32 = 0
	testRun(&counter, atomicAdd)
	testRun(&counter, lockAdd)
}

func testRun(counter *int32, f func(*int32)) {
	start := time.Now()
	wg := sync.WaitGroup{}
	num := 100
	wg.Add(num)
	run(num, counter, &wg, f)
	wg.Wait()
	end := time.Now()
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fmt.Printf("%v %v:%v\n", *counter, fn, end.Sub(start))
}

func run(num int, counter *int32, wg *sync.WaitGroup, f func(*int32)) {
	for i := 0; i < num; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				f(counter)
			}
			wg.Done()
		}()
	}
}

func atomicAdd(counter *int32) {
	atomic.AddInt32(counter, 1)
}

var lock = sync.Mutex{}

func lockAdd(counter *int32) {
	lock.Lock()
	defer lock.Unlock()
	*counter = *counter + 1
}
