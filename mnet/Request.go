package mnet

import (
	"github.com/MouseChannel/MouseChannelServer/face"
)

type Request struct {
	message []byte

	session face.ISession
}

func (request *Request) GetMessage() []byte {
	return request.message
}

func (request *Request) GetSession() face.ISession {
	return request.session
}

func NewRequest(message []byte, session face.ISession) Request {
	return Request{
		message: message,
		session: session,
	}

}
