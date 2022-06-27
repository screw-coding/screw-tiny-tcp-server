package main

import (
	"log"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	addr, err := net.ResolveTCPAddr("tcp", ":5200")
	if err != nil {
		log.Fatal(err)
		return
	}

	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Listening on port 5200")

	for {
		connection, err := tcp.Accept()
		if err != nil {
			log.Println("accept err:", err)
			return
		}

		log.Println("A toy_client connected.")
		go process(connection)
	}

}

func process(connection net.Conn) {
	defer func(connection net.Conn) {
		_ = connection.Close()
	}(connection)

	for {
		buffer := make([]byte, 1024)
		n, err := connection.Read(buffer)
		if err != nil {
			return
		}
		log.Println("received:", string(buffer[:n]), "from", connection.RemoteAddr())
		_, _ = connection.Write(buffer[:n])
	}
}
