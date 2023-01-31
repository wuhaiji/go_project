package context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
br:
	for true {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break br
		default:
			time.Sleep(time.Millisecond * 10)
		}
	}
	wg.Done()
}

func TestBackground(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(3)
	cancel()
	wg.Wait()
	fmt.Println("over")

}
func TestDeadline(t *testing.T) {
	d := time.Now().Add(5 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()
	select {
	case <-time.After(time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}

}

// test for whether it will block when use select sentence without default sub sentence
func TestSelect(t *testing.T) {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		ch1 <- 1
	}()
	select {
	case <-ch1:
	case <-ch2:
	}
	fmt.Println("done")
}

func TestCtxWithValue(t *testing.T) {

}
