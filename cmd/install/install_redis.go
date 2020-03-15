package install

import (
	"encoding/json"
	"fmt"
	software2 "hei/cmd/install/software"
	"hei/param"
	"hei/utils"
	"log"
)

type InstallerRedis struct {
	softWare *software2.SoftWare
}

func GetInstallerRedis(p param.Parameters) InstallerRedis {
	installer := InstallerRedis{}

	return installer
}

func (installer *InstallerRedis) Handle(p param.Parameters) (bool, string) {
	installer.softWare = &software2.SoftwareRedis
	var tag = ""
	if len(p.Args) > 0 {
		tag = p.Args[0]
	}
	//1.获取软件信息
	item, err := installer.softWare.GetItemByTag(tag)
	if !err {
		data, _ := json.Marshal(item)
		fmt.Printf("software Info : %s\n", data)
	} else {
		fmt.Printf("Can't find the software, something wrong.")
	}
	//2.下载文件
	download := utils.DownloadFile(item.Url, item.FileName)
	log.Printf("download ........:%t\n", download)

	//3.解压文件到目录
	utils.UnzipTarGz(item.FileName, "./")
	utils.MoveDocument("./redis*", param.GetConfig().AppPath+item.InstallFileName)
	//4.配置环境变量

	//5.添加配置文件

	return true, ""
}
