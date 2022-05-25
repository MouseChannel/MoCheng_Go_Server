package face

import "github.com/MouseChannel/MoChengServer/pb"

type IRoomState interface {
	Enter()
	Exit()
	Update(session ISession, mes *pb.PbMessage)
}
