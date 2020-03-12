package install

import "hei/param"

func  GetInstallerRedis(p param.Parameters) InstallerRedis {
	installer := InstallerRedis{}
	installer.items = make([]InstallItem,0)
	installer.items = append(installer.items, InstallItem{
		Tag:"aaa",
		Url:"bbb",
		FileName:"ccc",
	})
	//installer.config = []map[string] string {
	//	"tag": "jdk-8u202-linux-x64",
	//	"url": "12",
	//	"sex": "male",
	//}
	//installer.config = append(installer.config,
	//	)
	return installer
}
