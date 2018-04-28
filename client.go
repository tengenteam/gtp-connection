package gtp

import (
	"strings"
	"fmt"
)

type GTPClient struct {
	conn        *GTPConnection
	pro_version string
}

// 创建GTP命令
func NewGtpClient(conn *GTPConnection) *GTPClient {
	util := GTPClient{}
	util.conn = conn
	ver, err := conn.Exec("protocol_version")
	if err == nil {
		util.pro_version = ver
	} else {
		util.pro_version = "2"
	}
	return &util
}

// 判断命令是否支持
func (self GTPClient) KnowCommand(cmd string) (string, error) {
	return self.conn.Exec(cmd)
}

// 获取AI落子
func (self GTPClient) GenMove(color string) (string, error) {
	color = strings.ToUpper(color)
	command := "black"
	if color == "B" {
		command = "black"
	} else if color == "W" {
		command = "white"
	}
	if self.pro_version == "1" {
		command = "genmove_" + command
	} else {
		command = "genmove " + command
	}
	return self.conn.Exec(command)
}
// 人落子
func (self GTPClient) Move(color, coor string) (string, error) {
	color = strings.ToUpper(color)
	command := "black"
	if color == "B" {
		command = "black"
	} else if color == "W" {
		command = "white"
	}
	return self.conn.Exec(fmt.Sprintf("play %s %s", command, coor))
}

// 加载SGF文件
func (self GTPClient) LoadSgf(file string) (string, error) {
	command := fmt.Sprintf("loadsgf %s", file)
	return self.conn.Exec(command)
}
// 获取当前盘面形势判断
func (self GTPClient) FinalStatusList(cmd string) (string, error) {
	command := fmt.Sprintf("final_status_list %s", cmd)
	return self.conn.Exec(command)
}
// 设置AI级别
func (self GTPClient) SetLevel(seed int) (string, error) {
	command := fmt.Sprintf("level %d", seed)
	return self.conn.Exec(command)
}
// 设置AI随机数
func (self GTPClient) SetRandomSeed(seed int) (string, error) {
	command := fmt.Sprintf("set_random_seed %d", seed)
	return self.conn.Exec(command)
}
// 显示棋盘
func (self GTPClient) ShowBoard() (string, error) {
	return self.conn.Exec("showboard")
}
// 清空棋盘
func (self GTPClient) ClearBoard() (string, error) {
	return self.conn.Exec("clear_board")
}
//打印SGF
func (self GTPClient) PrintSgf() (string, error) {
	return self.conn.Exec("printsgf")
}
// 设置时间规则
func (self GTPClient) TimeSetting(baseTime, byoTime, byoStones int) (string, error) {
	return self.conn.Exec(fmt.Sprintf("time_settings %d %d %d", baseTime, byoTime, byoStones))
}
//获取结果
func (self GTPClient) FinalScore() (string, error) {
	return self.conn.Exec("final_score")
}
//退出
func (self GTPClient) Quit() (string, error) {
	return self.conn.Exec("Quit")
}
// 方式自定义命令
func (self GTPClient) SendCMD(cmd string) (string, error) {
	return self.conn.Exec(cmd)
}
