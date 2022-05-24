package roomSystem

import (
	"fmt"
	"server/face"
	"server/mnet/connectPool"
	"server/pb"
	"server/singleton"
)

type RoomMessageHandle struct {
	connectPool *connectPool.ConnectPool
}

func (messageHandle *RoomMessageHandle) Init() {
	messageHandle.connectPool = singleton.Singleton[connectPool.ConnectPool]()
}

func (messageHandle *RoomMessageHandle) Response(session face.ISession, message *pb.PbMessage) {
	// sid := session.GetSid()
	roomId := session.GetCurrentRoomId()
	if roomId == 0 {
		return

	}
	fmt.Println("Receive :", message)

	room := messageHandle.connectPool.GetRoom(roomId)
	if room != nil {

		room.GetCurrentState().Update(session, message)
	} else {
		fmt.Println("No Room")
	}

}
