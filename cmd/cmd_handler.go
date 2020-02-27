package cmd

import (
	help2 "hei/cmd/help"
	"hei/cmd/install"
	"hei/param"
)


type Handler struct {

}

func getHandler( p param.Parameters) HandlerInterface {
	if  p.Cmd == ""{
		return nil
	}
	switch p.Cmd {
	case "help":
		return help2.Help{}
	case "install":
		return install.Installer{}
	default:
		return nil
	}
}

func (handler * Handler) Handle(param param.Parameters)  {
	h  := getHandler(param)
	h.Handle(param)

}