package roomSystem

import (
	"fmt"

	"github.com/MouseChannel/MouseChannelServer/face"
	"github.com/MouseChannel/MouseChannelServer/pb"
)

type RoomStateLoadResource struct {
	room     face.IRoom
	percent  []int
	loadDone []bool
}

func NewRoomStateLoadResource(room face.IRoom, length int) face.IRoomState {
	state := &RoomStateLoadResource{
		room:     room,
		percent:  make([]int, length),
		loadDone: make([]bool, length),
	}
	return state
}

func (state *RoomStateLoadResource) Enter() {
	fmt.Println("进入load阶段")

	mes := pb.MakeRoomLoadCmd(state.room.GetSelectData())
	state.room.Broadcast(mes)

}
func (state *RoomStateLoadResource) Exit() {

}

func (state *RoomStateLoadResource) Update(session face.ISession, mes *pb.PbMessage) {
	index := state.room.GetPlayerIndex(session)
	loadPercent := mes.LoadPercent
	switch mes.CmdRoom{
	case pb.PbMessage_fightStart:
		state.loadDone[index] = true
		mes := pb.MakeRoomLoadData(100)
		state.room.Broadcast(mes)
		
		if state.CheckAllLoadDone() {
			state.room.ChangeRoomState(int(roomStateFight))
		}
	
	case pb.PbMessage_loadData:
		state.percent[index] = int(loadPercent)
		
		state.room.Broadcast(pb.Byte(mes))
	
}

}
func (state *RoomStateLoadResource) CheckAllLoadDone() bool {
	return true
}
