package install

import (
	"fmt"
	"hei/param"
	"hei/utils"
)

type InstallerRedis struct {
	config map[string] string
}


func  GetInstallerRedis(p param.Parameters) InstallerRedis {
	installer := InstallerRedis{}
	installer.config =map[string] string {
		"tag": "jdk-8u202-linux-x64",
		"url": "12",
		"sex": "male",
	}
	return installer
}


func (* InstallerRedis) Handle(p param.Parameters) (bool, string) {
	fmt.Println("Handle install redis command.......",p.Args)
	url :="https://repo.huaweicloud.com/redis/redis-5.0.7.tar.gz"
	utils.DownloadFile(url,"./")
	return true,""
}