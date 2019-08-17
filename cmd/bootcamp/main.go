package main

import (
	"fmt"
	"grab/internal/bootcamp"
)

func main() {
	go bootcamp.StartServer()

	_, _ = fmt.Scanln()

}
