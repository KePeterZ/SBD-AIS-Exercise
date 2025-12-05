package main

import (
	"exc8/client"
	"exc8/server"
	"time"
)

func main() {
	go func() {
		if err := server.StartGrpcServer(); err != nil {
			panic(err)
		}
	}()

	time.Sleep(1 * time.Second)
	c, err := client.NewGrpcClient()
	if err != nil {
		panic(err)
	}

	// Run client workflow
	if err := c.Run(); err != nil {
		panic(err)
	}
}
