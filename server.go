package main

import (
	"github.com/MouseChannel/MouseChannelServer/mnet"

	"fmt"
	// "github.com/MouseChannel/MouseChannelServer/singleton"
)

// mnet "github.com/MouseChannel/MouseChannelServer/Net"

func main() {

	mnet.ServerStartWork()
	// server.Serve()

}
func DoSomething() {
	fmt.Println("I do something")
}
