package main

import (
	"bufio"
	"log"
	"net"
	"os"
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
		for i := 0; i < 100; i++ {
			_, err = conn.Write([]byte(readString))
			if err != nil {
				log.Println("write err:", err)
				return
			}
		}

		buffer := make([]byte, 1024)
		size, err := conn.Read(buffer)
		if err != nil {
			return
		}

		if err != nil {
			log.Println("read err:", err)
			return
		}

		log.Println("received:", string(buffer[:size]), "from", conn.RemoteAddr())

	}

}
