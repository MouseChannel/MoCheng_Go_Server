package face

import (
	"github.com/MouseChannel/MCServer/pb"
)

type IMessageHandle interface {
	Init( )
	Response(session ISession, message *pb.PbMessage)
}
