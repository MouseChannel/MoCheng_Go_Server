package face

import "server/pb"

type IRoomState interface {
	Enter()
	Exit()
	Update(session ISession, mes *pb.PbMessage)
}
