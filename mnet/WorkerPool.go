package mnet

import (
	"fmt"
	"github.com/MouseChannel/MCServer/face"
	"github.com/MouseChannel/MCServer/mnet/connectPool"
	"github.com/MouseChannel/MCServer/mnet/matchSystem"
	"github.com/MouseChannel/MCServer/mnet/roomSystem"
	"github.com/MouseChannel/MCServer/pb"
	"github.com/MouseChannel/MCServer/singleton"

	"google.golang.org/protobuf/proto"
)

type WorkerPool struct {
	WorkerPoolSize uint32
	TaskQueue      []chan Request

	matchMessageHandle face.IMessageHandle
	roomMessageHandle  face.IMessageHandle
	connectPool        *connectPool.ConnectPool
}

func (workerPool *WorkerPool) Init() {
	fmt.Println("WorkerPool Init")

	workerPool.WorkerPoolSize = 10

	workerPool.TaskQueue = make([]chan Request, 10)
	workerPool.connectPool = singleton.Singleton[connectPool.ConnectPool]()
	workerPool.matchMessageHandle = singleton.Singleton[matchSystem.MatchMessageHandle]()
	// workerPool.matchMessageHandle.Init()
	workerPool.roomMessageHandle = singleton.Singleton[roomSystem.RoomMessageHandle]()
	// workerPool.roomMessageHandle.Init()

	workerPool.StartWorkerPool()

}

func (workerPool *WorkerPool) DoMessageHandler(request Request) {

	//测试下Pb能不能解码
	mes := &pb.PbMessage{}
	if err := proto.Unmarshal(request.GetMessage(), mes); err != nil {
		fmt.Println(err)
	}

	switch mes.Cmd {
	case pb.PbMessage_login:
		workerPool.ResponseLogin(request.GetSession().GetSid())

	case pb.PbMessage_match:
		workerPool.matchMessageHandle.Response(request.GetSession(), mes)
		// workerPool.matchMessageHandle.Response(request.GetSession(), mes)
	case pb.PbMessage_room:
		workerPool.roomMessageHandle.Response(request.GetSession(), mes)
		// workerPool.roomMessageHandle.Response(request.GetSession(), mes)
	case pb.PbMessage_chat:
		workerPool.ResponseTest(request.GetSession())
	}
}

func (workerPool *WorkerPool) ResponseLogin(sid uint32) {
	mes := pb.MakeLogin()
	workerPool.connectPool.GetSession(sid).SendMessage(mes)

	// workerPool.server.SendMessageToClient(sid, mes)

}

//test
func (workerPool *WorkerPool) ResponseTest(session face.ISession) {
	// mes := pb.Byte(mes1)

	// for sid := range workerPool.server.GetAllPlayer() {
	// 	workerPool.server.SendMessageToClient(sid, mes)
	// }

}

func (workerPool *WorkerPool) StartWorkerPool() {

	for i := 0; i < int(workerPool.WorkerPoolSize); i++ {
		workerPool.TaskQueue[i] = make(chan Request)
		go workerPool.StartOneWorker(i, workerPool.TaskQueue[i])
	}

}
func (workerPool *WorkerPool) StartOneWorker(workerID int, taskQueue chan Request) {

	// fmt.Println("WorkerId = ", workerID, "  Start")
	for {
		request := <-taskQueue
		workerPool.DoMessageHandler(request)
		fmt.Println(workerID, "work over")

	}
}
func (workerPool *WorkerPool) AddToTaskQueue(request Request) {

	workerID := request.GetSession().GetSid() % workerPool.WorkerPoolSize
	fmt.Println("AddTaskQueue  ", workerID)

	workerPool.TaskQueue[workerID] <- request

}

// func NewWorkerPool(_server face.IServer) *WorkerPool {
// 	return &WorkerPool{
// 		server:         _server,
// 		WorkerPoolSize: 10,
// 		TaskQueue:      make([]chan face.IRequest, 10),
// 		// roomMessageHandle: &room.RoomMessageHandle{},
// 	}

// }
