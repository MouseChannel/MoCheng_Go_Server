package main

import (
	"github.com/MouseChannel/MoChengServer/mnet"

	"fmt"
	// "github.com/MouseChannel/MoChengServer/singleton"
)

// mnet "github.com/MouseChannel/MoChengServer/Net"

func main() {

	mnet.ServerStartWork()
	// server.Serve()

}
func DoSomething() {
	fmt.Println("I do something")
}
