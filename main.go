package main

import (
	"github.com/FOXFOUND/Sylvia/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.RouterRegister(router)
	router.Run(":9000")
}
