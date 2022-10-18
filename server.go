package main

import (
	"github.com/MouseChannel/MoCheng_Go_Server/mnet"

	"fmt"
	// "github.com/MouseChannel/MoCheng_Go_Server/singleton"
)

// mnet "github.com/MouseChannel/MoCheng_Go_Server/Net"

func main() {

	mnet.Get_Server_Instance().Start()
	// server.Serve()

}
func DoSomething() {
	fmt.Println("I do something")
}
