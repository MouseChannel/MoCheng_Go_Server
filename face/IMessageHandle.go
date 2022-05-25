package face

import (
	"github.com/MouseChannel/MoChengServer/pb"
)

type IMessageHandle interface {
	Init()
	Response(session ISession, message *pb.PbMessage)
}
