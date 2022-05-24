package face

import "github.com/MouseChannel/MCServer/pb"

type IRoomState interface {
	Enter()
	Exit()
	Update(session ISession, mes *pb.PbMessage)
}
