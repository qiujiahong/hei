package cmd

import (
	"hei/param"
)


type HandlerInterface interface {
	Handle(param param.Parameters) (bool, string)
}

