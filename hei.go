package main

import (
	"hei/cmd"
	"hei/param"
)

// 环境: PRO DEV
var ENV="PRO"

func main() {
	//获取命令和参数
	parameters := param.Parameters{}
	parameters.Read()
	//执行参数
	handler := cmd.Handler{}
	handler.Handle(parameters)
}
