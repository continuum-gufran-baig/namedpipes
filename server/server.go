package server

import (
	"fmt"
	"time"

	"github.com/ContinuumLLC/pipes/client"
	"github.com/ContinuumLLC/platform-common-lib/src/namedpipes"
)

var oneMin = time.Second * 5

func CreateServer(pipeName string) {
	pipe, err := namedpipes.GetPipeServer().CreatePipe(pipeName, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err := pipe.Close()
		if err != nil {
			return
		}
	}()

	fmt.Println("pipe created!!")

	// accept incoming connection
	inc, err := pipe.Accept()
	if err != nil {
		fmt.Printf("Error Accepting %+v\n", err)
	}
	fmt.Println("accepted connection !!")

	// write back some data to conn
	client.WriteToPipeWithConn(inc)
}
