package param

import (
	"log"
	"os"
)

type Config struct {
	AppPath string
	Env string
}

var GConfig = Config{}

//获取配置文件
func GetConfig() Config  {
	runMode := os.Getenv("RUN_MODE")
	if runMode == "DEV"{
		log.Println("Run in dev mode")
		GConfig = Config{
			AppPath: "./install",
			//PRO DEV
			Env:     "DEV",
		}
	}else  {
		GConfig = Config{
			AppPath: "/apps/",
			//PRO DEV
			Env:     "PRO",
		}
	}
	return GConfig
}