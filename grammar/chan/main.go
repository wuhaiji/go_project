package main

import (
	"fmt"
	"grammar/channel"
	"time"
)

func main() {
	//var ch chan int
	//fmt.Printf("%v\n", ch)
	//ch = make(chan int)
	//fmt.Printf("%v\n", ch)

	//ch := make(chan int, 1)
	//go func() {
	//	//var x = <-ch
	//	//fmt.Printf("x:%v\n", x)
	//	for true {
	//		time.Sleep(time.Second)
	//	}
	//}()
	//ch <- 10
	//x := <-ch
	//fmt.Printf("x=%v", x)

	time.Sleep(time.Second)
	ch2 := channel.NewNoBufChannel[int]()
	go func() {
		msg := ch2.Recv()
		fmt.Printf("msg:%v\n", msg)
	}()

	ch2.Send(1)
}
