package matchSystem

import (
	"container/list"
	"fmt"
	"server/face"

	"server/mnet/roomSystem"
	"server/pb"
	"sync"
)

type MatchSystem struct {
	matchQueue list.List
	matchlen   int
	lock       *sync.Mutex
}

func (match *MatchSystem) Init() {

	match.matchlen = 4
	match.lock = new(sync.Mutex)

}

func (match *MatchSystem) UpdateMatchQueue(message *pb.PbMessage, session face.ISession) {

	switch message.CmdMatch {
	case pb.PbMessage_joinMatch:
		match.EnterMatchQueue(session)

	case pb.PbMessage_quitMatch:
		match.QuitMatchQueue(session)

	}

}

func (match *MatchSystem) EnterMatchQueue(session face.ISession) {
	fmt.Println("a player join match", session)
	match.lock.Lock()
	match.matchQueue.PushBack(session)

	match.lock.Unlock()

	mes := pb.MakeJoinMatch()
	session.SendMessage(mes)
	// match.server.SendMessageToClient(sid, mes)

	if match.matchQueue.Len() >= match.matchlen {
		match.GenerateNewRoom()
	}

}

func (match *MatchSystem) QuitMatchQueue(session face.ISession) {
	fmt.Println("a player quit match", session.GetSid())
	match.lock.Lock()
	for i := match.matchQueue.Front(); i != nil; i = i.Next() {
		if i.Value == session {
			match.matchQueue.Remove(i)
			fmt.Println("a player actually quit match", i, session.GetSid())

		}
	}
	match.lock.Unlock()
	mes := pb.MakeQuitMatch()
	session.SendMessage((mes))
	// match.server.SendMessageToClient(sid, mes)
}

func (match *MatchSystem) GenerateNewRoom() {

	match.lock.Lock()

	roomPlayerSession := make([]face.ISession, match.matchlen)
	for i := 0; i < match.matchlen; i++ {
		session := match.matchQueue.Front().Value.(face.ISession)
		roomPlayerSession[i] = session
		match.matchQueue.Remove(match.matchQueue.Front())
	}

	newRoom := roomSystem.NewRoom(roomPlayerSession)

	newRoom.Init()
	fmt.Println("newww rom")

	match.lock.Unlock()

}
