package mnet

import (
	"container/list"
	"fmt"

	"github.com/MouseChannel/MoCheng_Go_Server/face"

	"sync"

	"github.com/MouseChannel/MoCheng_Go_Server/pb"
)

type MatchManager struct {
	matchQueue list.List
	matchlen   int
	lock       *sync.Mutex
}

func (match *MatchManager) UpdateMatchQueue(message *pb.PbMessage, session face.ISession) {

	switch message.CmdMatch {
	case pb.PbMessage_joinMatch:
		match.EnterMatchQueue(session)

	case pb.PbMessage_quitMatch:
		match.QuitMatchQueue(session)

	}

}

func (match *MatchManager) EnterMatchQueue(session face.ISession) {
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

func (match *MatchManager) QuitMatchQueue(session face.ISession) {
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

func (match *MatchManager) GenerateNewRoom() {

	match.lock.Lock()

	roomPlayerSession := make([]face.ISession, match.matchlen)
	for i := 0; i < match.matchlen; i++ {
		session := match.matchQueue.Front().Value.(face.ISession)
		roomPlayerSession[i] = session
		match.matchQueue.Remove(match.matchQueue.Front())
	}

	SpawnRoom(roomPlayerSession)

	fmt.Println("newww rom")

	match.lock.Unlock()

}

func NewMatchSystem() face.IMatchSystem {
	return &MatchManager{
		matchlen: 4,
		lock:     new(sync.Mutex),
	}
}

var match_instance face.IMatchSystem
var match_once sync.Once

func Get_Match_Instance() face.IMatchSystem {
	match_once.Do(func() {
		match_instance = NewMatchSystem()
	})
	return match_instance
}
