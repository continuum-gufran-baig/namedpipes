//go:generate goversioninfo -64

package main

import (
	// "github.com/ContinuumLLC/pipes/client"
	"github.com/ContinuumLLC/pipes/server"
)

const pipe = `\\.\pipe\MyPipe`

func main() {
	server.CreateServer(pipe)
}
