package face

import "github.com/MouseChannel/MCServer/pb"

type IMatchSystem interface {
	Init()
	EnterMatchQueue(session ISession)
	QuitMatchQueue(session ISession)
	UpdateMatchQueue(message *pb.PbMessage, session ISession)
}
