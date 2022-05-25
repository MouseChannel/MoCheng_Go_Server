package face

import "github.com/MouseChannel/MouseChannelServer/pb"

type IMatchSystem interface {
	Init()
	EnterMatchQueue(session ISession)
	QuitMatchQueue(session ISession)
	UpdateMatchQueue(message *pb.PbMessage, session ISession)
}
