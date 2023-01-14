package _chan

import (
	"fmt"
	"grammar/sync/channel"
	"reflect"
	"sync"
	"testing"
	"time"
)

type ChInt = channel.Channel

func Producer() (ch channel.Channel) {
	ch = channel.NewBufChannel(2)
	go func() {
		for i := 0; i < 4; i++ {
			err := ch.Send(i)
			if err != nil {
				panic(err)
			}
		}
		//err := ch.Send("测试string")
		//if err != nil {
		//	panic(err)
		//}
		ch.Close()
	}()
	return ch
}
func Consumer(ch channel.Channel) int {
	var sum = 0
	ch.Range(func(data any) {
		d, ok := data.(int)
		if ok {
			sum += d
		} else {
			fmt.Printf("[error]receive a data is not int,%v", data)
		}
	})
	return sum
}

func Test1(t *testing.T) {
	ch := Producer()
	sum := Consumer(ch)
	fmt.Printf("sum = %v\n", sum)
	if sum != 6 {
		t.Error("sum error")
	}
}
func Test2(t *testing.T) {
	var ch chan int = make(chan int)
	go func() {
		x := <-ch
		fmt.Printf("%v\n", x)
	}()
	//ch <- 1
}
func Test3(t *testing.T) {
	//var ch = make(chan int, 1)
	//close(ch)
	//ch <- 1
	//chInt := channel.channel{}
	//fmt.Printf("%v\n", chInt)
}

func Test4(t *testing.T) {
	var ch = make(chan any)

	go func() {
		ch <- "123"
		ch <- 456
	}()

	var x = <-ch
	fmt.Printf("%v\n", x)
	x = <-ch
	fmt.Printf("%v\n", x)
}

type A struct {
}
type B struct {
}

func Test5(t *testing.T) {
	var a = A{}
	var b = B{}
	kindA := reflect.TypeOf(a).Kind()
	kindB := reflect.TypeOf(b).Kind()
	fmt.Printf("a kind：%v\n", kindA)
	fmt.Printf("a type：%v\n", reflect.TypeOf(a))
	fmt.Printf("b：%v\n", kindB)
	fmt.Printf("b type：%v\n", reflect.TypeOf(b))
	fmt.Printf("a==b = %v ", kindB == kindA)
	fmt.Printf("a==b = %v ", reflect.DeepEqual(kindB, kindA))
}

func Test6(t *testing.T) {
	var ch = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

func Test7(t *testing.T) {

	// demo1 通道误用导致的bug
	func() {
		wg := sync.WaitGroup{}

		ch := make(chan int, 10)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)

		wg.Add(3)
		for j := 0; j < 3; j++ {
			go func() {
				for {
					task, ok := <-ch
					if !ok {
						break
					}
					// 这里假设对接收的数据执行某些操作
					fmt.Println(task)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}()

}

func Test8(t *testing.T) {
	// demo2 通道误用导致的bug
	func() {
		ch := make(chan string, 1)
		go func() {
			// 这里假设执行一些耗时的操作
			time.Sleep(3 * time.Second)
			ch <- "job result"
			close(ch)
		}()

		select {
		case result := <-ch:
			fmt.Println(result)
		case <-time.After(time.Second): // 较小的超时时间
			return
		}
	}()
}
