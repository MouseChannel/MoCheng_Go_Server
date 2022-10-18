package MessageHandle

import (
	"github.com/MouseChannel/MoCheng_Go_Server/face"
	"github.com/MouseChannel/MoCheng_Go_Server/pb"
)

type MatchMessageHandle[T face.IMatchSystem] struct {
	matchSystem face.IMatchSystem
}

func (matchMessageHandle *MatchMessageHandle[T]) Response(session face.ISession, message *pb.PbMessage, data T) {

	data.UpdateMatchQueue(message, session)

}

func NewMatchMessageHandle() face.IMessageHandle[face.IMatchSystem] {
	return &MatchMessageHandle[face.IMatchSystem]{}
}
