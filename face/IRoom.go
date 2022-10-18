package face

// type RoomState int32

// const (
// 	roomStateConfirm      RoomState = 0
// 	roomStateSelect       RoomState = 1
// 	roomStateLoadResource RoomState = 2
// 	roomStateFight        RoomState = 3
// 	roomStateEnd          RoomState = 4
// 	roomStateNone         RoomState = 5
// )

type IRoom interface {
	Start()
	Delete()
	Broadcast(data []byte)
	// SendIndex()

	ChangeRoomState(newState int)
	ChangePlayersRoomId()
	// GetState() IRoomState
	GetCurrentState() IRoomState
	GetPlayerIndex(session ISession) int
	SetSelectData(selectArr []int)
	GetSelectData() []int

	GetRoomPlayerCount() int
	GetRoomId() uint32

	// UpdateConfirm(mes *pb.PbMessage)
}
