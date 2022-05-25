package roomSystem

import (
	"fmt"
	"time"

	"github.com/MouseChannel/MouseChannelServer/face"
	"github.com/MouseChannel/MouseChannelServer/pb"
)

// import "github.com/MouseChannel/MouseChannelServer/face"

type RoomStateConfirm struct {
	// *RoomStateBase
	room       face.IRoom
	confirmArr []bool
}

func NewRoomStateConfirm(room face.IRoom, length int) face.IRoomState {
	state := &RoomStateConfirm{
		room:       room,
		confirmArr: make([]bool, length),
	}
	return state
}

func (state *RoomStateConfirm) Enter() {
	for i := 0; i < state.room.GetRoomPlayerCount(); i++ {
		state.confirmArr[i] = false
	}

	state.room.ChangePlayersRoomId()

	mes := pb.MakeRoomConfirm()
	//向玩家发送进入房间命令，和房间成员信息

	state.room.Broadcast(mes)
	fmt.Println("Confirm")
	go state.CheckTimeLimit()

}
func (state *RoomStateConfirm) Dismiss() {

	mes := pb.MakeRoomDismiss()
	state.room.Broadcast(mes)
	state.room.Delete()

}
func (state *RoomStateConfirm) Exit() {

}

func (state *RoomStateConfirm) Update(sesssion face.ISession, mes *pb.PbMessage) {

	index := state.room.GetPlayerIndex(sesssion)
	state.confirmArr[index] = true

	if state.CheckAllConfirm() {
		// state.room.ChangeRoomState(int(roomStateSelect))

		// ----------------
		//测试用

		// mes := pb.MakeRoomLoadCmd([]int{})
		//向玩家发送进入房间命令，和房间成员信息

		// state.room.Broadcast(mes)

		state.room.ChangeRoomState(int(roomStateSelect))
		//=====================

	}

}
func (state *RoomStateConfirm) CheckTimeLimit() {
	<-time.After(time.Second * 99)
	if state == state.room.GetCurrentState() {
		fmt.Println("room confirm reachtime ")
		state.Dismiss()
	}

}

func (state *RoomStateConfirm) CheckAllConfirm() bool {

	for _, i := range state.confirmArr {
		if !i {
			return false
		}
	}
	fmt.Println("全部确认")
	return true

}
