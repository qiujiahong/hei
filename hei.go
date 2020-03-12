package main

import (
	"hei/cmd"
	"hei/param"
)



func main() {
	//获取命令和参数
	param.GetConfig()
	parameters := param.Parameters{}
	parameters.Read()
	//执行参数
	handler := cmd.Handler{}
	handler.Handle(parameters)
}
