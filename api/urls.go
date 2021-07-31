package api

import (
	"github.com/gin-gonic/gin"
)

type HTTP_STRUCT struct {
	HTTP_METHOD   string
	HTTP_FUNCTION func(*gin.Context)
}

var URL_PATTERNS = map[string]HTTP_STRUCT{
	"/healthz/": {
		"GET",
		healthz,
	},
	"/heartbeat/": {
		"POST",
		heartbeat,
	},
}

func RouterRegister(router *gin.Engine) {
	for URI, _struct := range URL_PATTERNS {
		switch _struct.HTTP_METHOD {
		case "GET":
			router.GET(URI, _struct.HTTP_FUNCTION)
		case "POST":
			router.POST(URI, _struct.HTTP_FUNCTION)
		}
	}
}
