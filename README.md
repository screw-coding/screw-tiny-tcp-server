# screw-tiny-tcp-server

go实现最简单的tcp客户端和服务端,实现带自定义包结构的tcp客户读和服务端

## 目录结构

```
├── go.mod
├── pack_client_server
│ ├── pack_client.go 带构建包和解析包的客户端
│ ├── pack_server.go 带构建包和解析包的服务端
│ └── protocol.go
├── README.md
└── toy_client_server
  ├── toy_client.go 最简单的TCP客户端
  └── toy_server.go 最简单的TCP服务端
```

## 说明

因为TCP是流式传输,可能出现客户端发送的内容,分成多段流到达服务端,如果服务端总是认为每次都是客户端一次发送的完整的内容,就会出现内容不完整的问题,因此需要在服务端处理检测内容是否完整的问题

toy_client.go运行后,在toy_server.go上可以看到,不一定是一个一个收的,日志里会显示不定长的消息在一起

而在pack_server中,每次都是一个完整的信息


