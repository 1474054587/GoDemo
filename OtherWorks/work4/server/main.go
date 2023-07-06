package main

import (
	"context"
	"fmt"
	"time"
)

// import (
//
//	"bufio"
//	"bytes"
//	"context"
//	"encoding/binary"
//	"fmt"
//	"io"
//	"net"
//
// )
//
//	func Decode(reader *bufio.Reader) (string, error) {
//		lengthByte, _ := reader.Peek(4)
//		lengthBuff := bytes.NewBuffer(lengthByte)
//		var length int32
//		err := binary.Read(lengthBuff, binary.LittleEndian, &length)
//		if err != nil {
//			return "", err
//		}
//		if int32(reader.Buffered()) < length+4 {
//			return "", err
//		}
//
//		pack := make([]byte, int(4+length))
//		_, err = reader.Read(pack)
//		if err != nil {
//			return "", err
//		}
//		return string(pack[4:]), nil
//	}
//
//	func process(conn net.Conn) {
//		defer conn.Close()
//		reader := bufio.NewReader(conn)
//		for {
//			msg, err := Decode(reader)
//			if err == io.EOF {
//				return
//			}
//			if err != nil {
//				fmt.Println("decode msg failed, err:", err)
//				return
//			}
//			fmt.Println(msg)
//		}
//	}
//
//	func main() {
//		listen, err := net.Listen("tcp", "127.0.0.1:8080")
//		if err != nil {
//			fmt.Println("listen failed, err:", err)
//			return
//		}
//		defer listen.Close()
//		for {
//			conn, err := listen.Accept()
//			if err != nil {
//				fmt.Println("accept failed, err:", err)
//				continue
//			}
//			go process(conn)
//		}
//	}

func main() {
	// fmt.Println("start:", time.Now())
	// father, cancel1 := context.WithTimeout(context.TODO(), time.Second*5)
	// defer cancel1()
	// son, cancel2 := context.WithTimeout(father, time.Second*10)
	// defer cancel2()
	// deadline, ok := son.Deadline()
	// fmt.Printf("son deadline:%v ok:%v\n", deadline, ok)
	ctx_timeout, _ := context.WithTimeout(context.TODO(), time.Second*1)
	select {
	case <-ctx_timeout.Done():
		fmt.Println("ctx timeout, err:", ctx_timeout.Err())
	}
	ctx_cancel, cancel := context.WithCancel(context.TODO())
	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()
	select {
	case <-ctx_cancel.Done():
		fmt.Println("ctx cancel err:", ctx_cancel.Err())
	}
}
