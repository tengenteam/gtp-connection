# GTP-Connection

`gtp` 是一个基于exec.Cmd(golang)实现的Go Text Protocol的GTP连接器，我们可以轻松的使用此连接器快速与围棋AI进行对弈。

## Installation

    go get -u github.com/tengenteam/gtp-connection

## 简单的Gnugo实现!

```go
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
```

## GTPConnection

如何创建一个连接器，和执行命令

```go
NewGTPConnection(cmd string, args ...string) (*GTPConnection, error)
Exec(cmd string) (string, error)
```

## Client
Client是一个基于连接器封装的与AI对弈的常用指令

```go
NewGtpClient(conn *GTPConnection) *GTPClient
KnowCommand(cmd string) (string, error)
GenMove(color string) (string, error)
Move(color, coor string) (string, error)
LoadSgf(file string) (string, error)
FinalStatusList(cmd string) (string, error)
SetLevel(seed int) (string, error) 
SetRandomSeed(seed int) (string, error)
ShowBoard() (string, error) 
ClearBoard() (string, error) 
PrintSgf() (string, error)
TimeSetting(baseTime, byoTime, byoStones int) (string, error) 
FinalScore() (string, error)
Quit() (string, error) 
SendCMD(cmd string) (string, error)
```