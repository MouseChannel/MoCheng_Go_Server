package mnet

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"github.com/MouseChannel/MoChengServer/face"
	"github.com/MouseChannel/MoChengServer/singleton"

	"github.com/MouseChannel/MoChengServer/mnet/matchSystem"

	"github.com/MouseChannel/MoChengServer/mnet/connectPool"

	"time"

	"github.com/xtaci/kcp-go/v5"
)

type Server struct {
	IP             string
	UDPIP          string
	connectionPool *connectPool.ConnectPool
	workerPool     face.IWorkerPool
	matchSystem    face.IMatchSystem
}

func (server *Server) Start() {
	go server.PrintLogo()

	go server.ListenUDP()

	go server.ListenKCP()
}

//UDP用于用户初次连接，分配一个KCPsession给他
func (server *Server) ListenUDP() {

	address, _ := net.ResolveUDPAddr("udp", server.UDPIP)
	// remoteaddress, _ := net.ResolveUDPAddr("udp", "127.0.0.1:7777")

	conn, err := net.ListenUDP("udp", address)

	if err != nil {
		fmt.Println("listen UDP failed!!")
	}
	for {

		buf := make([]byte, 32)
		_, remoteAddress, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("receive UDP failed!!")
		}
		// fmt.Println(remoteAddress)
		// fmt.Println(buf)
		//这里的连接是本地（只有一个）的连接吗？如果是的话还需要把udp连接关闭，暂时还不知道怎么写
		go func() {
			if binary.BigEndian.Uint32(buf[:4]) == 0 {

				//某个客户端申请连接,分配一个sid给它
				buf := make([]byte, 4)
				sid := server.connectionPool.GenerateUniqueSessionID()
				fmt.Println("UDPid", sid)
				binary.BigEndian.PutUint32(buf, sid)
				// fmt.Println("unique sid ", buf)
				// fmt.Println(append([]byte{0, 0, 0, 0}, []byte(buf)...))
				conn.WriteToUDP(append([]byte{0, 0, 0, 0}, []byte(buf)...), remoteAddress)

			}
		}()
	}
}

//KCP就是建立连接后的正常业务
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
		server.AddSession(conn)

	}
}

func (server *Server) Init() {
	fmt.Println("github.com/MouseChannel/MoChengServer Init")
	server.IP = "0.0.0.0:6666"
	server.UDPIP = "0.0.0.0:7777"
	server.connectionPool = singleton.Singleton[connectPool.ConnectPool]()
	// server.connectionPool.Init()

	singleton.Singleton[matchSystem.MatchSystem]().Init()

}

func (server *Server) Serve() {

	// server.Init()
	server.Start()

	select {}

}
func (server *Server) AddSession(conn *kcp.UDPSession) {
	sid := conn.GetConv()

	newSession := NewSession(conn, sid)
	server.connectionPool.AddSession(sid, newSession)
	newSession.Start()

	fmt.Println("a new session connect ")
}
func (server *Server) RemoveSession(sid uint32) {
	server.connectionPool.DeleteSession(sid)
	// delete(server.sessionMap, sid)
}
func (server *Server) GetSession(sid uint32) face.ISession {

	return server.connectionPool.GetSession(sid)

}

func (server *Server) Stop() {

}

//测试用

func (server *Server) PrintLogo() {

	file, err := os.Open("./logo.txt")
	if err != nil {
		fmt.Print("err")
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

func ServerStartWork() {
	singleton.Singleton[Server]().Serve()
}
