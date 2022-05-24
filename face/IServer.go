package face

import "github.com/xtaci/kcp-go/v5"

type IServer interface {
	Start()
	Stop()
	Serve()

	GetRoom(roomId uint32) IRoom
	AddRoom(roomId uint32, room IRoom)
	RemoveRoom(roomId uint32)

	GetSession(sid uint32) ISession
	AddSession(conn *kcp.UDPSession)
	RemoveSession(sid uint32)

	HandleMessage(request IRequest)
	SendMessageToClient(sid uint32, data []byte)

	GetMatchSystem() IMatchSystem

	GetAllPlayer() map[uint32]ISession
}
