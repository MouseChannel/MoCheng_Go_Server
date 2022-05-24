package face

import "server/pb"

type IMatchSystem interface {
	Init()
	EnterMatchQueue(session ISession)
	QuitMatchQueue(session ISession)
	UpdateMatchQueue(message *pb.PbMessage, session ISession)
}
