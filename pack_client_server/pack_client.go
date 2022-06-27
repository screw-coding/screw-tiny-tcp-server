package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"screw-tiny-tcp-server/pack_client_server/my_packet"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5200")
	if err != nil {
		log.Fatal("dial err:", err)
		return
	}
	input := bufio.NewReader(os.Stdin)
	for {
		readString, err := input.ReadString('\n')
		if err != nil {
			return
		}
		readString = strings.TrimSpace(readString)

		if readString == "quit" {
			log.Println("quit")
			return
		}

		for i := 0; i < 10; i++ {
			_, err = conn.Write(my_packet.Pack([]byte(readString)))
			if err != nil {
				log.Println("write err:", err)
				return
			}
		}

		readerChannel := make(chan []byte, 16)
		go clientRead(readerChannel)
		go receiveAndUnpack(conn, readerChannel)

	}

}

func receiveAndUnpack(conn net.Conn, readerChannel chan []byte) {
	buffer := make([]byte, 1024)
	tmpBuffer := make([]byte, 0)
	for {
		size, err := conn.Read(buffer)
		if err != nil {
			return
		}
		if err != nil {
			log.Println("read err:", err)
			return
		}
		my_packet.Unpack(append(tmpBuffer, buffer[:size]...), readerChannel)
	}
}

func clientRead(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			log.Println("客户端收到数据:", string(data))
		}
	}
}
