package main

import (
	"bufio"
	"fmt"
	"grammar/net/socket/custom_protocol"
	"io"
	"net"
)

func main() {
	tcp, err := net.Listen("tcp", "127.0.0.1:9090")
	fmt.Println("listening tcp 127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	for true {
		conn, err2 := tcp.Accept()
		if err2 != nil {
			fmt.Printf("listen failed ,err:%v\n", err2)
			continue
		}
		fmt.Printf("收到tcp连接,remote address: %v\n", conn.RemoteAddr())
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
		if err := recover(); err != nil {
			fmt.Printf("网路异常 err:%v\n", err)
		}
	}()
	for {
		fmt.Print("")
		reader := bufio.NewReader(conn)
		s, err := custom_protocol.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", s)
		msg := "" + s
		pkg, _ := custom_protocol.Encode(&msg)
		_, err = conn.Write(pkg)
		if err != nil {
			panic(err)
		}
	}
}
