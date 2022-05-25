package face

import (
	"github.com/MouseChannel/MouseChannelServer/pb"
)

type IMessageHandle interface {
	Init()
	Response(session ISession, message *pb.PbMessage)
}
