package face

import "github.com/MouseChannel/MouseChannelServer/pb"

type IRoomState interface {
	Enter()
	Exit()
	Update(session ISession, mes *pb.PbMessage)
}
