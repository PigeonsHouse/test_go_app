package main

import (
	"fmt"
	"test_go_app/db"
	"test_go_app/routers"
	"test_go_app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	utils.LoadEnv()
	db.InitDB()
	routers.InitRouter(api)
	api.Run(fmt.Sprintf(":%s", utils.ApiPort))
}
