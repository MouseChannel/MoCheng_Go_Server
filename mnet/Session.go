package mnet

import (
	"fmt"
	"net"

	"time"

	"github.com/MouseChannel/MoChengServer/face"
	"github.com/MouseChannel/MoChengServer/singleton"

	"github.com/xtaci/kcp-go/v5"
)

type Session struct {
	sid uint32

	roomId uint32

	kcpSession *kcp.UDPSession
	address    string
	isAlive    chan bool

	messageChan chan []byte

	// server face.IServer
}

func (session *Session) CheckAlive() {
	defer session.Stop()
	for {
		select {
		case <-session.isAlive:

		case <-time.After(time.Second * 9999):
			fmt.Println("Session die ")
			return
		}
	}
}

func (session *Session) StartReader() {
	fmt.Println("Session Start Read")
	// defer session.Stop()
	for {

		buf := make([]byte, 4096)

		n, err := session.kcpSession.Read(buf)
		//只截取有效数据部分

		if err != nil {
			fmt.Println("session read data failed!!")
			return
		}

		request := NewRequest(
			buf[:n],
			session)

		// mes := &pb.PbMessage{}
		// if err := proto.Unmarshal(request.GetMessage(), mes); err != nil {
		// 	fmt.Println(err)
		// }
		singleton.Singleton[WorkerPool]().AddToTaskQueue(request)
		// session.server.HandleMessage(request)

		session.isAlive <- true
	}
}

func (session *Session) SendMessage(data []byte) {

	session.messageChan <- data
}

func (session *Session) StartWriter() {
	for {
		data := <-session.messageChan
		if _, err := session.kcpSession.Write(data); err != nil {
			fmt.Println("Send Data error:, ", err, " Conn Writer exit")
			return
		}
	}

}

func (session *Session) Start() {
	go session.StartReader()

	go session.StartWriter()
	go session.CheckAlive()
}

func (session *Session) Stop() {
	session.kcpSession.Close()
	fmt.Println("session :", session.GetSid(), "STOP")
	singleton.Singleton[Server]().RemoveSession(session.sid)
	// session.server.RemoveSession(session.sid)

	//关闭该链接全部管道
	close(session.isAlive)
	close(session.messageChan)
}

func (session *Session) ChangeRoomId(roomId uint32) {
	session.roomId = roomId
}
func (session *Session) GetCurrentRoomId() uint32 {
	return session.roomId
}

func (session *Session) GetConnection() net.Conn {
	return session.kcpSession
}

func (session *Session) GetSid() uint32 {
	return session.sid
}

func (session *Session) GetRemoteAddress() string {
	return session.address
}
func NewSession(conn *kcp.UDPSession, sid uint32) face.ISession {
	session := &Session{
		sid:         sid,
		roomId:      0, // 该玩家属于哪个间
		kcpSession:  conn,
		isAlive:     make(chan bool),
		messageChan: make(chan []byte),
		// server:      server,
	}
	return session
}
