package main

import (
	"log"
	"net"
	"runtime"
	"screw-tiny-tcp-server/pack_client_server/my_packet"
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

		log.Println("A client connected.")
		go process(connection)
	}

}

func process(connection net.Conn) {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	go read(readerChannel, connection)
	buffer := make([]byte, 1024)
	for {
		n, err := connection.Read(buffer)
		if err != nil {
			log.Println(connection.RemoteAddr(), "connection error:", err)
			return
		}
		tmpBuffer = my_packet.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
		log.Println("read:", tmpBuffer)
	}
}

func read(readerChannel chan []byte, conn net.Conn) {
	for {
		select {
		case data := <-readerChannel:
			log.Println("服务端收到数据:", string(data))
			_, err := conn.Write(my_packet.Pack(append([]byte("服务端收到了且返回处理过的数据:"), data...)))
			if err != nil {
				log.Println("返回数据失败:", err)
			}
		}
	}
}
