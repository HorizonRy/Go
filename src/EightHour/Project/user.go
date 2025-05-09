package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// 创建新用户的 API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// 为每个用户启动一个 goroutine
	go user.ListenMessage()

	return user
}

// 监听当前 User channel，有消息时转发给客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C

		write, err := u.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("conn.Write err:", err, "write:", write)
			return
		}
	}
}
