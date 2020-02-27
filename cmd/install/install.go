package install

import (
	"fmt"
	"hei/param"
)


type Installer struct {

}




func (Installer) Handle(p param.Parameters) (bool, string) {
	if len(p.Args) < 1 {
		return false,"ERROR: need more parameters"
	}
	switch p.Args[0]  {
	case "java":
		java := GetInstallerJava(p)
		p.Args = p.Args[1:]
		java.Handle(p)
	case "redis":
		redis := GetInstallerRedis(p)
		p.Args = p.Args[1:]
		redis.Handle(p)

	default:
		fmt.Println("default install")
	}
	return true,""
}