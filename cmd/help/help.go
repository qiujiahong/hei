package help

import (
	"fmt"
	"hei/param"
)

type Help struct {

}
func (Help) Handle(p param.Parameters) (bool, string) {
	fmt.Println("Handle help command.")

	return true,""
}