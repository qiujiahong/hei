package param

import (
	"log"
	"os"
)

type Config struct {
	AppPath string
	Env string
	AppUser string
	AppGroup string
}
//配置
var gConfig = Config{}

//初始化配置
func InitConfig()  {
	runMode := os.Getenv("RUN_MODE")
	if runMode == "DEV"{
		log.Println("Run in dev mode")
		gConfig = Config{
			AppPath: "./install",
			//PRO DEV
			Env:     "DEV",
			AppGroup: nil,
			AppUser: nil,
		}
	}else  {
		gConfig = Config{
			AppPath: "/apps/",
			//PRO DEV
			Env:     "PRO",
			AppGroup: "appgroup",
			AppUser: "appuser",
		}
	}
}

//获取配置
func GetConfig() Config  {
	return gConfig
}