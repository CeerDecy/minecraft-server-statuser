package router

import (
	"github.com/gin-gonic/gin"

	"minecraft-server-statuser/internal/handle"
)

func WithRouter(router *gin.Engine) {
	router.GET("/status", handle.StatusHandle)
}
