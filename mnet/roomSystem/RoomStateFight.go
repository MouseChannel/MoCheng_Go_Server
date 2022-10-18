package roomSystem

import (
	"github.com/MouseChannel/MoCheng_Go_Server/face"
	"github.com/MouseChannel/MoCheng_Go_Server/pb"
	"github.com/MouseChannel/MoCheng_Go_Server/timer/timer"
	"github.com/MouseChannel/MoCheng_Go_Server/timer/timerface"
)

type RoomStateFight struct {
	room          face.IRoom
	mytimerticker timerface.ITimer
	fightOpArr    []*pb.FightMessage
	frameId       int
	endArr        []bool
}

func NewRoomStateFight(room face.IRoom, length int) face.IRoomState {
	state := &RoomStateFight{
		room:          room,
		mytimerticker: &timer.MyTimer{},
		fightOpArr:    []*pb.FightMessage{},
		frameId:       0,
		endArr:        make([]bool, length),
	}
	return state
}

func (state *RoomStateFight) Enter() {

	mes := pb.MakeFightStartCmd()
	state.room.Broadcast(mes)

	state.mytimerticker.AddTickTimerTask(66, state.SyncLogicFrame)

}
func (state *RoomStateFight) Exit() {

}

func (state *RoomStateFight) Update(sesssion face.ISession, mes *pb.PbMessage) {

	state.fightOpArr = append(state.fightOpArr, mes.SendFightMessage)

}

func (state *RoomStateFight) SyncLogicFrame() {
	state.frameId++

	mes := pb.MakeFightData(state.frameId, state.fightOpArr)

	state.room.Broadcast(mes)
	if len(state.fightOpArr) > 0 {
		state.fightOpArr = []*pb.FightMessage{}
	}

}

func (state *RoomStateFight) ResponseFightEnd(playerIndex int) {
	state.endArr[playerIndex] = true
	if state.CheckAllFightDone() {
		state.room.ChangeRoomState(int(roomStateEnd))
	}
}
func (state *RoomStateFight) CheckAllFightDone() bool {
	for _, i := range state.endArr {
		if !i {
			return false
		}
	}
	return true
}
