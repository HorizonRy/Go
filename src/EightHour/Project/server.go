package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	onlineUsersMap     map[string]*User
	onlineUsersMapLock sync.RWMutex

	// 消息广播 channel
	broadcastMsgChan chan string
}

// 创建一个 server 实例的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:               ip,
		Port:             port,
		onlineUsersMap:   make(map[string]*User),
		broadcastMsgChan: make(chan string),
	}

	return server
}

// 监听用户发送的消息
func (s *Server) ListenUserMsg() {
	for {
		msg := <-s.broadcastMsgChan
		s.onlineUsersMapLock.Lock()
		for _, user := range s.onlineUsersMap {
			user.C <- msg
		}
		s.onlineUsersMapLock.Unlock()
	}
}

// 用户上线广播
func (s *Server) Broadcast(newOnliner *User, msg string) {
	sendMsg := "[" + newOnliner.Addr + "]" + newOnliner.Name + ":" + msg

	s.broadcastMsgChan <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	// 处理用户上线
	// 创建一个用户
	user := NewUser(conn)

	// 将用户加入 onlineUsersMap
	s.onlineUsersMapLock.Lock()
	s.onlineUsersMap[user.Name] = user
	s.onlineUsersMapLock.Unlock()

	// 广播上线消息
	s.Broadcast(user, "已上线")

	// 当前 handler 阻塞，并创建监听用户消息的 Goroutine
	select {}
}

// 启动服务的接口
func (s *Server) Start() {
	// socket 监听
	listener, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%d", s.Ip, s.Port),
	)
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("listener.Close err:", err)
		}
	}(listener)

	// 启动监听用户消息的 Goroutine
	go s.ListenUserMsg()

	for {
		// 接收数据
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		// 处理请求
		go s.Handler(conn)
	}

}
