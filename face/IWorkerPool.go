package face

//判断此时玩家在对局里还是在大厅里
type IWorkerPool interface {
 
	DoMessageHandler(request IRequest) //处理该信息
	Start ()                  //启动工作池
	AddToTaskQueue(request IRequest)   //将消息添加到 某个工作池队列
 
}
