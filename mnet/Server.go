package mnet

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/MouseChannel/MoCheng_Go_Server/face"
	"github.com/xtaci/kcp-go/v5"

	"sync"
	"time"
)

type Server struct {
	IP    string
	UDPIP string
}

func (server *Server) Start() {
	go server.PrintLogo()

	go server.ListenUDP()

	go server.ListenKCP()
}

// UDP用于用户初次连接，分配一个KCPsession给他
func (server *Server) ListenUDP() {

	address, _ := net.ResolveUDPAddr("udp", server.UDPIP)

	conn, err := net.ListenUDP("udp", address)

	if err != nil {
		fmt.Println("listen UDP failed!!")
	}
	buf := make([]byte, 4)
	for {

		_, remoteAddress, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("receive UDP failed!!")
		}

		go func() {
			if binary.BigEndian.Uint32(buf) == 0 {

				//某个客户端申请连接,分配一个sid给它
				// buf := make([]byte, 4)
				sid := Get_ConnectPool_Instance().GenerateUniqueSessionID()
				fmt.Println("UDPid", sid)
				binary.BigEndian.PutUint32(buf, sid)
				// fmt.Println("unique sid ", buf)
				// fmt.Println(append([]byte{0, 0, 0, 0}, []byte(buf)...))
				conn.WriteToUDP(append([]byte{0, 0, 0, 0}, []byte(buf)...), remoteAddress)

			}
		}()
	}
}

// KCP就是建立连接后的正常业务
func (server *Server) ListenKCP() {
	//开启服务器端口

	kcplisten, err := kcp.ListenWithOptions(server.IP, nil, 0, 0)
	if err != nil {
		fmt.Println("kcp.Listen failed!!")
	}
	for {
		//监听是否有新的客户端连接
		conn, err := kcplisten.AcceptKCP()
		if err != nil {
			fmt.Println("accept conn failed!!")
			continue
		}
		Get_ConnectPool_Instance().AddSession(conn)

	}
}

func (server *Server) Stop() {

}

func (server *Server) PrintLogo() {

	file, err := os.Open("github.com/MouseChannel/MoCheng_Go_Server/logo.txt")
	if err != nil {
		// fmt.Print("Print Logo failed But still fine")
		fmt.Print("█▄ ▄█ ▄▀▄ ▄▀▀ █▄█ ██▀ █▄ █ ▄▀\n█ ▀ █ ▀▄▀ ▀▄▄ █ █ █▄▄ █ ▀█ ▀▄█\n█ ▄▀▀   █▀▄ ██▀ ▄▀▄ █▀▄ ▀▄▀\n█ ▄██   █▀▄ █▄▄ █▀█ █▄▀  █ \n")

		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		fmt.Print(string(line))
		if err == io.EOF {
			break
		}

		time.Sleep(time.Duration(100) * time.Millisecond)
	}

}

var server_instance face.IServer
var server_once sync.Once

func Get_Server_Instance() face.IServer {
	server_once.Do(func() {
		server_instance = NewServer()
	})
	return server_instance
}

func NewServer() face.IServer {
	return &Server{
		IP:    "0.0.0.0:6666",
		UDPIP: "0.0.0.0:7777",
	}

}
