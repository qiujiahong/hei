package param

import (
	"fmt"
	"os"
)

type Parameters struct {
	//命令字
	Cmd string
	//参数
	Args []string
}

func (param* Parameters) Read()  {
	if len(os.Args) <= 1 {
		param.Cmd = "help"
		return
	}
	param.Cmd = os.Args[1]
	param.Args = os.Args[2:]
	fmt.Printf("parameters: ")
	for i, v := range os.Args {
		//fmt.Printf("args[%v]=%v\n", i, v)
		if i != 0 {
			fmt.Printf("%v ",v)
		}
	}
	fmt.Println()
}