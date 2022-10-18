package face

import "github.com/MouseChannel/MoCheng_Go_Server/pb"

type IMatchSystem interface {
	EnterMatchQueue(session ISession)
	QuitMatchQueue(session ISession)
	UpdateMatchQueue(message *pb.PbMessage, session ISession)
}
