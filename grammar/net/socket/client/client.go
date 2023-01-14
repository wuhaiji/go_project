package main

import (
	"bufio"
	"errors"
	"fmt"
	"grammar/net/socket/custom_protocol"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	TestClient()
}
func TestClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	fmt.Println("dial tcp 127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn) // 关闭连接

re:
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
			break re
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
		fmt.Printf("user input err:%v\n", err)
		return "", err
	}
	inputStr = strings.Trim(inputStr, "\r\n")
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
