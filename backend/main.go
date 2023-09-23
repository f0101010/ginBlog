package main

import (
	"ginBlog/model"
	"ginBlog/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()
}
