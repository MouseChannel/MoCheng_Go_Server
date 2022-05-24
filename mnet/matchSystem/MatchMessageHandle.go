package matchSystem

import (
	"github.com/MouseChannel/MCServer/face"
	"github.com/MouseChannel/MCServer/pb"
	"github.com/MouseChannel/MCServer/singleton"
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
