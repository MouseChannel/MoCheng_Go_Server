package mnet

import (
	"github.com/MouseChannel/MoCheng_Go_Server/face"
	"github.com/MouseChannel/MoCheng_Go_Server/mnet/roomSystem"
)

func OnRoomSpawn(room face.IRoom) {
	Get_ConnectPool_Instance().AddRoom(room)
}
func OnRoomDelete(room face.IRoom) {
	Get_ConnectPool_Instance().DeleteRoom(room.GetRoomId())
}

func SpawnRoom(playerSessions []face.ISession) face.IRoom {
	newRoom := roomSystem.NewRoom(playerSessions, OnRoomSpawn, OnRoomDelete)
	newRoom.Start()

	return newRoom
}
