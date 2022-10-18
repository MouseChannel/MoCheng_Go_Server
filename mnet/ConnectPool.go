package mnet

import (
	"fmt"
	"sync"

	"github.com/MouseChannel/MoCheng_Go_Server/face"
	"github.com/xtaci/kcp-go/v5"
)

type ConnectionManager struct {
	sessionMap    map[uint32]face.ISession //存放所有玩家的连接
	roomMap       map[uint32]face.IRoom    //存放所有房间
	mux           sync.Mutex
	NewSessionSid uint32
}

func (connectionManager *ConnectionManager) AddSession(conn *kcp.UDPSession) {
	sid := conn.GetConv()
	newSession := NewSession(conn, sid)
	newSession.Start()
	connectionManager.sessionMap[sid] = newSession
	fmt.Println("a new session connect ")

}

func (connectionManager *ConnectionManager) DeleteSession(sid uint32) {
	delete(connectionManager.sessionMap, sid)
}

func (connectionManager *ConnectionManager) GetSession(sid uint32) face.ISession {
	if session, ok := connectionManager.sessionMap[sid]; ok {
		return session
	}
	return nil
}

func (connectionManager *ConnectionManager) GenerateUniqueSessionID() uint32 {
	connectionManager.mux.Lock()
	for {
		_, ok := connectionManager.sessionMap[connectionManager.NewSessionSid]
		if !ok {
			break
		}
		connectionManager.NewSessionSid++
	}
	connectionManager.sessionMap[connectionManager.NewSessionSid] = nil

	connectionManager.mux.Unlock()
	return connectionManager.NewSessionSid
}

//Room

func (connectionManager *ConnectionManager) AddRoom(room face.IRoom) {
	connectionManager.roomMap[room.GetRoomId()] = room
}

func (connectionManager *ConnectionManager) DeleteRoom(roomId uint32) {
	delete(connectionManager.roomMap, roomId)
}

func (connectionManager *ConnectionManager) GetRoom(roomId uint32) face.IRoom {
	if room, ok := connectionManager.roomMap[roomId]; ok {
		return room
	}
	return nil
}

// func (connectionManager *ConnectPool) Init() {}

func NewConnectionManager() face.IConnection {
	return &ConnectionManager{
		sessionMap:    make(map[uint32]face.ISession), //存放所有玩家的连接
		roomMap:       make(map[uint32]face.IRoom),    //存放所有房间
		mux:           *new(sync.Mutex),
		NewSessionSid: 1,
	}
}

var connection_instance face.IConnection
var connection_once sync.Once

func Get_ConnectPool_Instance() face.IConnection {
	connection_once.Do(func() {
		connection_instance = NewConnectionManager()
	})
	return connection_instance
}
