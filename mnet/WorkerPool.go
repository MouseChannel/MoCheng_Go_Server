package mnet

import (
	"fmt"
	"sync"

	"github.com/MouseChannel/MoCheng_Go_Server/face"
	"github.com/MouseChannel/MoCheng_Go_Server/mnet/MessageHandle"

	"github.com/MouseChannel/MoCheng_Go_Server/pb"

	"google.golang.org/protobuf/proto"
)

type WorkerPool struct {
	WorkerPoolSize uint32
	TaskQueue      []chan face.IRequest
	matchHandle    face.IMessageHandle[face.IMatchSystem]
	roomHandle     face.IMessageHandle[face.IRoom]
}

func (workerPool *WorkerPool) DoMessageHandler(request face.IRequest) {

	//测试下Pb能不能解码
	mes := &pb.PbMessage{}
	if err := proto.Unmarshal(request.GetMessage(), mes); err != nil {
		fmt.Println(err)
	}
	// a :=  Get_ConnectPool_Instance().GetSession( request.GetSid())
	// m:= MessageHandle.NewMatchMessageHandle()
	// m.Response(a,mes,Get_Match_Instance())

	switch mes.Cmd {
	case pb.PbMessage_login:
		workerPool.ResponseLogin(request.GetSid())

	case pb.PbMessage_match:
		workerPool.matchHandle.Response(request.GetSession(), mes, Get_Match_Instance())

	case pb.PbMessage_room:
		workerPool.roomHandle.Response(request.GetSession(), mes, request.GetRoom())
		// workerPool.roomMessageHandle.Response(request.GetSession(), mes)
	case pb.PbMessage_chat:
		workerPool.ResponseTest(request.GetSession())
	}
}

func (workerPool *WorkerPool) ResponseLogin(sid uint32) {
	mes := pb.MakeLogin()
	Get_ConnectPool_Instance().GetSession(sid).SendMessage(mes)

	// workerPool.server.SendMessageToClient(sid, mes)

}

// test
func (workerPool *WorkerPool) ResponseTest(session face.ISession) {
	// mes := pb.Byte(mes1)

	// for sid := range workerPool.server.GetAllPlayer() {
	// 	workerPool.server.SendMessageToClient(sid, mes)
	// }

}

func (workerPool *WorkerPool) Start() {

	for i := 0; i < int(workerPool.WorkerPoolSize); i++ {
		workerPool.TaskQueue[i] = make(chan face.IRequest)
		go workerPool.StartOneWorker(i, workerPool.TaskQueue[i])
	}

}
func (workerPool *WorkerPool) StartOneWorker(workerID int, taskQueue chan face.IRequest) {

	// fmt.Println("WorkerId = ", workerID, "  Start")
	for {
		request := <-taskQueue
		workerPool.DoMessageHandler(request)
		fmt.Println(workerID, "work over")

	}
}
func (workerPool *WorkerPool) AddToTaskQueue(request face.IRequest) {

	workerID := request.GetSid() % workerPool.WorkerPoolSize
	fmt.Println("AddTaskQueue  ", workerID)

	workerPool.TaskQueue[workerID] <- request

}

func NewWorkerPool() face.IWorkerPool {
	return &WorkerPool{

		WorkerPoolSize: 10,

		TaskQueue:   make([]chan face.IRequest, 10),
		matchHandle: MessageHandle.NewMatchMessageHandle(),
		roomHandle:  MessageHandle.NewRoomMessageHandle(),
	}

}

var worker_pool_instance face.IWorkerPool
var worker_pool_once sync.Once

func Get_WorkerPool_Instance() face.IWorkerPool {
	worker_pool_once.Do(func() {
		worker_pool_instance = NewWorkerPool()
	})
	return worker_pool_instance
}
