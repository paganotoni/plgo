package main

import (
	"github.com/paganotoni/plgo/web/server"
)

func main() {
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
