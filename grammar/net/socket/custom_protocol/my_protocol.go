package custom_protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

var byteOrder = binary.LittleEndian

func Encode(msg *string) (str []byte, err error) {
	length := int32(len(*msg))
	pkg := &bytes.Buffer{}

	err = binary.Write(pkg, byteOrder, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, byteOrder, []byte(*msg))
	if err != nil {
		return nil, err
	}

	by := pkg.Bytes()
	//fmt.Printf("encode %v\n", by)
	return by, nil
}

func Decode(connReader *bufio.Reader) (string, error) {
	// 读取消息的长度
	// 读取前4个字节的数据
	lengthByte, _ := connReader.Peek(4)
	var pkgLength int32
	lengthBuff := bytes.NewBuffer(lengthByte)
	err := binary.Read(lengthBuff, byteOrder, &pkgLength)

	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	length := pkgLength + 4
	if int32(connReader.Buffered()) < length {
		return "", nil
	}
	buf := make([]byte, length)
	n, err := connReader.Read(buf)

	if err != nil {
		return "", err
	}
	b := buf[:n]
	//fmt.Printf("decode %v\n", b)
	resp := string(b)
	return resp, nil
}
