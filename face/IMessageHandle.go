package face

import (
	"github.com/MouseChannel/MoCheng_Go_Server/pb"
)

type IMessageHandle[T any] interface {
	Response(session ISession, message *pb.PbMessage, data T)
}
