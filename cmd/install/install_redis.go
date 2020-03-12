package install

import (
	"encoding/json"
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
	var tag = ""
	if len(p.Args) > 0{
		tag = p.Args[0]
	}
    item, err := installer.softWare.GetItemByTag(tag)
    if !err{
		data, _ := json.Marshal(item)
		fmt.Printf("software Info : %s\n",data)
	}else {
		fmt.Printf("Can't find the software, something wrong.")
	}

	utils.DownloadFile(item.Url,item.FileName)
	return true,""
}