package connectPool

import (
	"fmt"
	"sync"

	"github.com/MouseChannel/MouseChannelServer/face"
)

type ConnectPool struct {
	sessionMap    map[uint32]face.ISession //存放所有玩家的连接
	roomMap       map[uint32]face.IRoom    //存放所有房间
	mux           sync.Mutex
	NewSessionSid uint32
}

//Single Session
func (connectPool *ConnectPool) AddSession(sid uint32, newSession face.ISession) {
	connectPool.sessionMap[sid] = newSession
}
func (connectPool *ConnectPool) DeleteSession(sid uint32) {
	delete(connectPool.sessionMap, sid)
}

func (connectPool *ConnectPool) GetSession(sid uint32) face.ISession {
	if session, ok := connectPool.sessionMap[sid]; ok {
		return session
	}
	return nil
}

func (connectPool *ConnectPool) GenerateUniqueSessionID() uint32 {
	connectPool.mux.Lock()
	for {
		_, ok := connectPool.sessionMap[connectPool.NewSessionSid]
		if !ok {
			break
		}
		connectPool.NewSessionSid++
	}
	connectPool.sessionMap[connectPool.NewSessionSid] = nil

	connectPool.mux.Unlock()
	return connectPool.NewSessionSid
}

//Room

func (connectPool *ConnectPool) AddRoom(roomId uint32, room face.IRoom) {
	connectPool.roomMap[roomId] = room
}

func (connectPool *ConnectPool) DeleteRoom(roomId uint32) {
	delete(connectPool.roomMap, roomId)
}

func (connectPool *ConnectPool) GetRoom(roomId uint32) face.IRoom {
	if room, ok := connectPool.roomMap[roomId]; ok {
		return room
	}
	return nil
}
func (connectPool *ConnectPool) Init() {
	connectPool.sessionMap = make(map[uint32]face.ISession) //存放所有玩家连接
	connectPool.roomMap = make(map[uint32]face.IRoom)       //存放所房间
	connectPool.mux = *new(sync.Mutex)
	connectPool.NewSessionSid = 1
	fmt.Println("Connect Pool Init")

}

// func (connectPool *ConnectPool) Init() {}

func NewConnectPool() *ConnectPool {
	return &ConnectPool{
		sessionMap:    make(map[uint32]face.ISession), //存放所有玩家的连接
		roomMap:       make(map[uint32]face.IRoom),    //存放所有房间
		mux:           *new(sync.Mutex),
		NewSessionSid: 1,
	}
}
