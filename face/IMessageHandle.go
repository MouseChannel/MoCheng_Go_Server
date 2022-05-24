package face

import (
	"server/pb"
)

type IMessageHandle interface {
	Init( )
	Response(session ISession, message *pb.PbMessage)
}
