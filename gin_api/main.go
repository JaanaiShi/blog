package main

import (
	"Blog/model"
	"Blog/routes"
)

func main() {
	// 引用数据库
	model.InitDb()

	routes.InitRouter()
}
