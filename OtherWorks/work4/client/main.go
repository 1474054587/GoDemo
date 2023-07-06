package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		message := fmt.Sprintf("id=%d,time=%v", i, time.Now())
		data, err := Encode(message)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		if _, err = conn.Write(data); err != nil {
			fmt.Println("write conn failed, err:", err)
			return
		}
	}
}
