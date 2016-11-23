package main

import (
	"fmt"
	"github.com/dist-ribut-us/ipc"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify the port to send to")
		return
	}

	msg := "Hello"
	if len(os.Args) > 2 {
		msg = strings.Join(os.Args[2:], " ")
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	srv, err := ipc.RunNew(0)
	if err != nil {
		panic(err)
	}

	srv.Send([]byte(msg), port)
}
