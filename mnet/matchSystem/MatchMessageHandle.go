package matchSystem

import (
	"github.com/MouseChannel/MoChengServer/face"
	"github.com/MouseChannel/MoChengServer/pb"
	"github.com/MouseChannel/MoChengServer/singleton"
)

type MatchMessageHandle struct {
	matchSystem face.IMatchSystem
}

func (messageHandle *MatchMessageHandle) Init( ) {

}

func (matchMessageHandle *MatchMessageHandle) Response(session face.ISession, message *pb.PbMessage) {
	// sid := session.GetSid()
	singleton.Singleton[MatchSystem]().UpdateMatchQueue(message, session)
	// matchMessageHandle.matchSystem
}

func NewMatchMessageHandle(matchSystem face.IMatchSystem) face.IMessageHandle {
	return &MatchMessageHandle{
		matchSystem: matchSystem,
	}
}
