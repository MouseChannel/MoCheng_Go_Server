package roomSystem

import (
	"fmt"
	"github.com/MouseChannel/MoChengServer/face"
	"github.com/MouseChannel/MoChengServer/pb"
	"time"
)

type RoomStateSelect struct {
	room      face.IRoom
	readyArr  []bool
	selectArr []int
}

func NewRoomStateSelect(room face.IRoom, length int) face.IRoomState {
	state := &RoomStateSelect{
		room:      room,
		readyArr:  make([]bool, length),
		selectArr: make([]int, length),
	}
	return state
}

func (state *RoomStateSelect) Enter() {
	for i := 0; i < state.room.GetRoomPlayerCount(); i++ {
		state.readyArr[i] = false
	}
	mes := pb.MakeRoomSelect()
	state.room.Broadcast(mes)
	fmt.Println("RoomSelect")
	go state.CheckTimeLimit()
}

func (state *RoomStateSelect) Dismiss() {

	mes := pb.MakeRoomDismiss()
	state.room.Broadcast(mes)
	state.room.Delete()

}
func (state *RoomStateSelect) Exit() {

}

func (state *RoomStateSelect) CheckTimeLimit() {
	<-time.After(time.Second * 999)
	fmt.Println("room confirm reachtime ")
	state.Dismiss()

}

func (state *RoomStateSelect) Update(session face.ISession, mes *pb.PbMessage) {
	index := state.room.GetPlayerIndex(session)

	state.readyArr[index] = mes.SelectData.IsReady
	state.selectArr[index] = int(mes.SelectData.Faction)
	// if !mes.SelectData.IsReady{

	// }

	// selectdata := pb.MakeRoomSelectData(mes,)
	mes.RoomIndex = int32(index)
	state.room.Broadcast(pb.Byte(mes))

	// mes := pb.MakeSelectData()
	if state.CheckAllSelectDone() {
		state.room.SetSelectData(state.selectArr)
		state.room.ChangeRoomState(int(roomStateLoadResource))
	}
}

func (state *RoomStateSelect) CheckAllSelectDone() bool {
	for _, i := range state.readyArr {
		if !i {
			return false
		}
	}
	return true

}
