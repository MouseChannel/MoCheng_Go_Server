package pb

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func Byte(mes *PbMessage) []byte {
	if data, err := proto.Marshal(mes); err != nil {
		fmt.Println(err)
	} else {
		return data
	}
	return nil
}

func MakeLogin() []byte {

	mes := &PbMessage{
		Cmd:  PbMessage_login,
		Name: "login",
	}
	return Byte(mes)

}
func MakeJoinMatch() []byte {
	mes := &PbMessage{
		Cmd:      PbMessage_match,
		CmdMatch: PbMessage_joinMatch,
	}
	return Byte(mes)
}
func MakeQuitMatch() []byte {
	mes := &PbMessage{
		Cmd:      PbMessage_match,
		CmdMatch: PbMessage_quitMatch,
	}
	return Byte(mes)
}

//server 发给client: Cmd,CmdRoom: PbMessage_confirm

//client 发给server: Cmd,CmdRoom: PbMessage_confirm, sid

func MakeRoomConfirm() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_confirm,
	}
	return Byte(mes)
}
 

func MakeRoomDismiss() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_dismissed,
	}
	return Byte(mes)

}

// func MakeSelectData( int) []byte{
// 	mes := &PbMessage{

// 	}
// }

func MakeRoomSelect() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_select,
	}
	return Byte(mes)
}

// func MakeRoomSelectData(singleMes *PbMessage, index int) []byte{
 
// 	newSelectData := SelectData{
// 		Index:         int32(index),
// 		PlayerName:    "",
// 		Faction:       0,
// 		IsReady:       false,
// 	}
// 	mes := &PbMessage{
// 		Cmd:     PbMessage_room,
// 		CmdRoom: PbMessage_selectDate,
// 		SelectData: &newSelectData,

// 	}
// 	return Byte(mes)

// }

func MakeRoomLoadCmd(roomSelectData []int) []byte {
	newRoomSelectData := []*SelectData{}
	for _, i := range roomSelectData {
		newRoomSelectData = append(newRoomSelectData, &SelectData{
			PlayerName: string(i),
			Faction:    int32(i)})
	}
	mes := &PbMessage{
		Cmd:            PbMessage_room,
		CmdRoom:        PbMessage_load,
		RoomSelectData: newRoomSelectData,
	}

	return Byte(mes)
}
func MakeRoomLoadData(loadPercent int32) []byte {
	mes := &PbMessage{
		Cmd:         PbMessage_room,
		CmdRoom:     PbMessage_loadData,
		LoadPercent: loadPercent,
	}
	return Byte(mes)
}
 


func MakeFightStartCmd() []byte {
	mes := &PbMessage{
		Cmd:     PbMessage_room,
		CmdRoom: PbMessage_fightStart,
	}
	return Byte(mes)
}

func MakeFightData(frameId int, fightMessage []*FightMessage) []byte {
	mes := &PbMessage{
		Cmd:          PbMessage_fight,
		FrameId:      int32(frameId),
		FightMessage: fightMessage,
	}
	return Byte(mes)
}
func MakeFightEnd() []byte {
	mes := &PbMessage{
		Cmd: PbMessage_fightEnd,
	}
	return Byte(mes)
}
