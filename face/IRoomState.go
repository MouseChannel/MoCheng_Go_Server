package face

import "github.com/MouseChannel/MoCheng_Go_Server/pb"

type IRoomState interface {
	Enter()
	Exit()
	Update(session ISession, mes *pb.PbMessage)
}
