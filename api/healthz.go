package api

import "github.com/gin-gonic/gin"

type HeartBeatJson struct {
	Id     string `json:"id" binding:"required`
	Status string `json:"status" binding:"required"`
}

// The Sylvia Master Healthz Check
func healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func heartbeat(c *gin.Context) {
	var _json HeartBeatJson
	c.BindJSON(&_json)
	// Your Business Here
	// Your Business End
	c.JSON(200, gin.H{
		"msg":  "success",
		"code": 200,
		"id":   _json.Id,
	})
}
