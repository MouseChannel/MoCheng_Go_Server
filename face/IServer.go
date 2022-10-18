package face

// import "github.com/xtaci/kcp-go/v5"

type IServer interface {
	Start()
	Stop()
	ListenUDP()
	ListenKCP()

	// GetRoom(roomId uint32) IRoom
	// AddRoom(roomId uint32, room IRoom)
	// RemoveRoom(roomId uint32)

	 
	 
	 

	// HandleMessage(request IRequest)
	// SendMessageToClient(sid uint32, data []byte)

	// GetMatchSystem() IMatchSystem

	// GetAllPlayer() map[uint32]ISession
}
