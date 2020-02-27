package install

import (
	"fmt"
	"hei/param"
	"hei/utils"
)

type InstallerJava struct {
	config map[string] string
}

//https://mirrors.huaweicloud.com/java/jdk/8u202-b08/jdk-8u202-linux-x64.tar.gz
func  GetInstallerJava(p param.Parameters) InstallerJava {
	installer := InstallerJava{}
	installer.config =map[string] string {
		"tag": "jdk-8u202-linux-x64",
		"url": "12",
		"sex": "male",
	}
	return installer
}


func (* InstallerJava) Handle(p param.Parameters) (bool, string) {
	fmt.Println("Handle install java command.......",p.Args)
	url :="https://mirrors.huaweicloud.com/java/jdk/8u202-b08/jdk-8u202-linux-x64.tar.gz"
	utils.DownloadFile(url,"./")
	return true,""
}