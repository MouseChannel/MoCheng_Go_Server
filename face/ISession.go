package face

import "net"

// "net"

type ISession interface {
	Start()
	Stop()
	GetConnection() net.Conn
	GetSid() uint32
	ChangeRoomId(roomId uint32)
	GetCurrentRoomId() uint32

	GetRemoteAddress() string
	SendMessage(data []byte)
}
