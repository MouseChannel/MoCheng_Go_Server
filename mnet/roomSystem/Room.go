package roomSystem

import (
	"github.com/MouseChannel/MoCheng_Go_Server/face"

	"sync"
)

type RoomState int

const (
	roomStateConfirm      RoomState = 0
	roomStateSelect       RoomState = 1
	roomStateLoadResource RoomState = 2
	roomStateFight        RoomState = 3
	roomStateEnd          RoomState = 4
	roomStateNone         RoomState = 5
)

type SelectData struct {
	selectId int
	isReady  bool
}

// a single room
type Room struct {
	stateMap     map[RoomState]face.IRoomState
	currentState RoomState
	roomId       uint32

	playerSessions []face.ISession
	selectArr      []int

	lock     sync.Mutex
	onSpawn  func(face.IRoom)
	onDelete func(face.IRoom)
}

func (room *Room) Start() {
	room.onSpawn(room)
	length := len(room.playerSessions)
	//Add StateMap
	room.stateMap[roomStateConfirm] = NewRoomStateConfirm(room, length)
	room.stateMap[roomStateSelect] = NewRoomStateSelect(room, length)
	room.stateMap[roomStateLoadResource] = NewRoomStateLoadResource(room, length)
	room.stateMap[roomStateFight] = NewRoomStateFight(room, length)
	// room.stateMap[roomStateEnd] = &RoomStateEnd{room}

	room.ChangeRoomState(int(roomStateConfirm))

}

func (room *Room) Delete() {
	room.onDelete(room)

}

func (room *Room) Broadcast(data []byte) {
	if data != nil {
		for _, session := range room.playerSessions {
			session.SendMessage(data)
			// room.server.SendMessageToClient(sid, data)
		}
	}
}

func (room *Room) ChangeRoomState(newState int) {

	if int(room.currentState) != newState {
		room.stateMap[room.currentState].Exit()

	}

	room.currentState = RoomState(newState)

	room.stateMap[room.currentState].Enter()

}
func (room *Room) ChangePlayersRoomId() {
	for _, session := range room.playerSessions {
		session.ChangeRoomId(room.roomId)
		// room.server.GetSession(i).ChangeRoomId(room.roomId)
	}
}

func (room *Room) SendLoadResource() {

}

func (room *Room) SendFightStart() {

}

func (room *Room) GetRoomId() uint32 {
	return room.roomId
}
func (room *Room) GetRoomPlayerCount() int {
	return len(room.playerSessions)
}
func (room *Room) GetPlayerIndex(session face.ISession) int {
	for index, i := range room.playerSessions {
		if i == session {
			return index
		}
	}
	return -1
}

func (room *Room) SetSelectData(selectArr []int) {
	room.selectArr = selectArr
}
func (room *Room) GetSelectData() []int {
	return room.selectArr
}
func (room *Room) GetCurrentState() face.IRoomState {
	return room.stateMap[room.currentState]
}

func NewRoom(playerSessions []face.ISession, OnSpawn func(face.IRoom), OnDelete func(face.IRoom)) face.IRoom {

	newRoom := &Room{

		stateMap:       make(map[RoomState]face.IRoomState),
		currentState:   roomStateConfirm,
		roomId:         playerSessions[0].GetSid(),
		playerSessions: playerSessions,
		lock:           *new(sync.Mutex),
		onSpawn:        OnSpawn,
		onDelete:       OnDelete,
	}

	return newRoom
}
