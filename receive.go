package main

import (
	"fmt"
	"log"
	"net"

	protobuf_test "github.com/pj-rione/protobuf-test/proto/pb_gen"
	"google.golang.org/protobuf/proto"
)

func main() {

	// receive data from UDP
	// UDP 受信処理
	conn, err := net.ListenPacket("udp", "127.0.0.1:1234")

	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		fmt.Println("Error reading UDP data:", err)
		return
	}

	fmt.Println("UDP data received successfully")
	fmt.Printf("Data received: %x \n", buffer[:n])

	p1 := &protobuf_test.DetectObjects{}
	err = proto.Unmarshal(buffer[:n], p1)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}

	fmt.Println(p1)

	log.Println(p1.GetObject())
}
