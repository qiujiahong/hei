package software

var SoftwareRedis = SoftWare{
	Name: "redis",
	Versions:  []Item{
		Item{
			Tag : "5.0.7",
			Url :"https://repo.huaweicloud.com/redis/redis-5.0.7.tar.gz",
			FileName: "redis-5.0.7.tar.gz",
			Default:true,
		},
		Item{
			Tag : "5.0.6",
			Url :"https://repo.huaweicloud.com/redis/redis-5.0.6.tar.gz",
			FileName: "redis-5.0.6.tar.gz",
			Default:true,
		},
	},
}