package roomSystem

import (
	"github.com/MouseChannel/MouseChannelServer/face"
	"github.com/MouseChannel/MouseChannelServer/pb"
)

type RoomStateEnd struct {
	room face.IRoom
}

func (state *RoomStateEnd) Enter() {
	mes := pb.MakeFightEnd()
	state.room.Broadcast(mes)
	state.Exit()

}
func (state *RoomStateEnd) Exit() {
	state.room.Delete()
}

func (state *RoomStateEnd) Update(sid uint32, mes *pb.PbMessage) {

}
