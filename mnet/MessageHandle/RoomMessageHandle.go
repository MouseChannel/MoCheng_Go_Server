package MessageHandle

import (
	"fmt"

	"github.com/MouseChannel/MoCheng_Go_Server/face"

	"github.com/MouseChannel/MoCheng_Go_Server/pb"
)

type RoomMessageHandle[T face.IRoom] struct {
}

func (messageHandle *RoomMessageHandle[T]) Response(session face.ISession, message *pb.PbMessage, room T) {
	// sid := session.GetSid()
	roomId := session.GetCurrentRoomId()
	if roomId == 0 {
		return

	}

	fmt.Println("Receive :", message)

	room.GetCurrentState().Update(session, message)

}

func NewRoomMessageHandle() face.IMessageHandle[face.IRoom] {
	return &RoomMessageHandle[face.IRoom]{}
}
