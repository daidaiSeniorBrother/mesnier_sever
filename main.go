package main

import (
	"mesnier/confs"
	"mesnier/routers"
)

func main() {
	// 初始化各项配置
	confs.SetUp()
	//初始化路由
	r := routers.InitRouter()
	_ = r.Run("0.0.0.0:8765")
}
