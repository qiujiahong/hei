package install

import (
	"fmt"
	"hei/param"
	"hei/utils"
)

type InstallerRedis struct {
	config  [] map[string] string
	items []InstallItem
}




func (* InstallerRedis) Handle(p param.Parameters) (bool, string) {
	fmt.Println("Handle install redis command.......",p.Args)
	url :="https://repo.huaweicloud.com/redis/redis-5.0.7.tar.gz"
	utils.DownloadFile(url,"redis-5.0.7.tar.gz")
	return true,""
}