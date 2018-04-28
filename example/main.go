package main

import (
	"fmt"
	"os"
	"github.com/tengenteam/gtp-connection"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
func main() {
	conn, err := gtp.NewGTPConnection("./gnugo_mac", "--mode", "gtp")
	checkError(err)
	client:=gtp.NewGtpClient(conn)
	move,err:=client.GenMove("B")
	fmt.Println(move,err)
	move,err=client.GenMove("W")
	fmt.Println(move,err)
	board,err:=client.ShowBoard()
	fmt.Println(board,err)
}
