package main

import (
	"main/controllers"
	"main/models"
)

func main() {
	models.ConnectDatabase()
	controllers.GetRankList(0)
}
