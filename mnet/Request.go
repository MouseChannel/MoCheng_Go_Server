package mnet

import "github.com/MouseChannel/MoCheng_Go_Server/face"

type Request struct {
	message []byte
	sid     uint32
}

func (request *Request) GetMessage() []byte {
	return request.message
}

func (request *Request) GetSid() uint32 {
	return request.sid
}
func (request *Request) GetSession() face.ISession {
	return Get_ConnectPool_Instance().GetSession(request.sid)
}
func (request *Request) GetRoom() face.IRoom {
	return Get_ConnectPool_Instance().GetRoom(request.sid)
}
func NewRequest(message []byte, sid uint32) *Request {
	return &Request{
		message: message,
		sid:     sid,
	}

}
