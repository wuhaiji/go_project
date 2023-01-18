package main

import (
	"bufio"
	"errors"
	"fmt"
	"grammar/go_id"
	"grammar/net/socket/custom_protocol"
	"grammar/util"
	"io"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	//
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go TestClient()
	}
	wg.Wait()
	fmt.Printf("goroutine count %v", runtime.NumGoroutine())
	time.Sleep(time.Second * 4)
}
func TestClient() {
	fmt.Printf("[%v]start client\n", go_id.GetGoroutineID())
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	fmt.Println("dial tcp 127.0.0.1:9090")
	if err != nil {
		panic(err)
	}

	fmt.Printf("local address %v \n", conn.LocalAddr())
	// 关闭连接
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

br:
	for {
		inputStr, ok := input()
		switch ok {
		case nil:
			write(conn, inputStr)
			msg, err := read(conn)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(msg)
		case io.EOF:
			break br
		default:
			panic(err)
		}
	}
}

var quitError = errors.New("user quit client")

func input() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('\n')
	if err == io.EOF {
		return "", err
	}
	if err != nil {
		util.Printlnf("user input err:%v", err)
		return "", err
	}
	inputStr = strings.Trim(inputStr, "\r\n")
	fmt.Printf("[%v]你的输入是:%v\n", go_id.GetGoroutineID(), inputStr)
	if strings.ToLower(inputStr) == "q" {
		fmt.Println("exit...")
		return "", quitError
	}
	return inputStr, nil
}

func read(conn net.Conn) (string, error) {
	connReader := bufio.NewReader(conn)
	msg, err := custom_protocol.Decode(connReader)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func write(conn net.Conn, inputStr string) {
	pkg, _ := custom_protocol.Encode(&inputStr)
	_, err := conn.Write(pkg)
	if err != nil {
		panic(err)
	}
}

//
