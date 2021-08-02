package PCodeLimbo

import (
	"fmt"
	"github.com/xStrixU/PCodeLimbo/player"
	"net"
)

var PlayerManager player.PlayerManager

func Start() {
	listener, err := net.Listen("tcp", "localhost:25566")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("listening!")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		player := player.NewPlayer(conn)
		go player.Start()
	}
}