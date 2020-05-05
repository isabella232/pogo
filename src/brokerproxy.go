package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"./pb"
	"google.golang.org/protobuf/proto"
	// "google.golang.org/protobuf/proto"
)

// Server initializes
func main() {
	port := AssignString(os.Getenv("PulsarPort"), ":6650")
	fmt.Printf("starting log server on port %s\n", port)

	// TODO: add TLS implementation
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	conn, err := lis.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	log.Println("client connected")
	defer conn.Close()

	data := make([]byte, 4096)
	cmd := pb.BaseCommand{}
	length, err := conn.Read(data)
	if err != nil {
		log.Printf("read data error")
		panic(err)
	}
	// log.Println(string(data))
	if err := proto.Unmarshal(data, &cmd); err != nil {
		log.Printf("unmarshal data error")
		log.Fatal(err)
	}

	log.Printf("Hello world sent, length %d bytes", length)
}
