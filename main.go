package main

import (
	"github.com/oyrabbit/dream-fun-admin/model"
	"github.com/oyrabbit/dream-fun-admin/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 引入路由组件
	routes.InitRouter()

}
