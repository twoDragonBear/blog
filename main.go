package main

import (
	"blog/model"
	"blog/router"
)

func main() {
	//init db
	model.InitDb()
	router.InitRouter()
}