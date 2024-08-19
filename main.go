package main

import (
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/pj-rione/protobuf-test/proto/pb_gen"
)

// 送信側プログラム
func main() {
	hash := "fugafuga"
	p := &pb_gen.DetectObjects{
		SrcId: "hogehoge",
		Object: []*pb_gen.DetectObjects_Object{
			&pb_gen.DetectObjects_Object{
				Type:       1,
				Confidence: 0.8,
				Position: &pb_gen.Vector2D{
					X: 1.0,
					Y: 2.0,
				},
				Length: 3.0,
			},
		},
		Hash: &hash,
	}

	//p.String() を hexで表示
	fmt.Println(p.String())
	fmt.Printf("p.String() = %x\n", p.String())

	body, _ := proto.Marshal(p)

	fmt.Println("============================")

	//body を hexで表示
	fmt.Printf("body = %x\n", body)

	//body を UDP で送信する
	// UDP送信処理
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(body)
	if err != nil {
		fmt.Println("Error sending UDP data:", err)
		return
	}

	fmt.Println("UDP data sent successfully")

}
