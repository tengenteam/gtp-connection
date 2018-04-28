package gtp

import (
	"os/exec"
	"io"
	"fmt"
	"bufio"
	"strings"
	"errors"
)

type GTPConnection struct {
	cmd     *exec.Cmd
	infile  io.WriteCloser
	outfile io.ReadCloser
}

// NewGTPConnection 创建GTP连接器
func NewGTPConnection(cmd string, args ...string) (*GTPConnection, error) {
	conn := GTPConnection{}
	conn.cmd = exec.Command(cmd, args...)
	inf, err := conn.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	outf, err := conn.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	conn.infile = inf
	conn.outfile = outf
	conn.cmd.Start()
	go func() {
		conn.cmd.Wait()
	}()
	return &conn, nil
}
// Exec 执行交互命令
func (g GTPConnection) Exec(cmd string) (string, error) {
	g.infile.Write([]byte(fmt.Sprintf("%s \n", cmd)))
	reader := bufio.NewReader(g.outfile)
	result := ""
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if line == "\n" {
			break
		} else if line == "\r\n" {
			break
		}
		result += line
	}
	res := strings.Split(result, "")
	l := len(res)
	if res[l-1] == "\n" {
		result = strings.Join(res[:l-1], "")
	}
	if len(result) == 0 {
		return "", errors.New("ERROR length=0")
	}
	if res[0]=="="{
		return strings.Join(res[2:],""),nil
	}
	if res[0]=="?"{
		return "",errors.New(strings.Join(res[2:],""))
	}
	return "",errors.New("ERROR: Unrecognized answer: " + result)
}

// Close 释放GTP资源
func (g GTPConnection) Close() {
	g.infile.Close()
	g.outfile.Close()
}
