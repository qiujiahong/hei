package install

import (
	"fmt"
	software2 "hei/cmd/install/software"
	"hei/param"
	"hei/utils"
)

type InstallerRedis struct {
	softWare *  software2.SoftWare
}

func  GetInstallerRedis(p param.Parameters) InstallerRedis {
	installer := InstallerRedis{}

	return installer
}


func ( installer * InstallerRedis) Handle(p param.Parameters) (bool, string) {
	installer.softWare = & software2.SoftwareRedis

	fmt.Println(installer.softWare.Versions)

	fmt.Println("Handle install redis command.......",p.Args)
	url :="https://repo.huaweicloud.com/redis/redis-5.0.7.tar.gz"
	utils.DownloadFile(url,"redis-5.0.7.tar.gz")
	return true,""
}