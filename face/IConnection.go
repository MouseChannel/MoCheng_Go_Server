package face

import "github.com/xtaci/kcp-go/v5"

type IConnection interface {
	AddSession(conn *kcp.UDPSession)
	DeleteSession(sid uint32)
	GetSession(sid uint32) ISession
	GenerateUniqueSessionID() uint32
	AddRoom( room IRoom)
	DeleteRoom(roomId uint32)
	GetRoom(roomId uint32) IRoom
}
