package main

import (
	"github.com/gin-gonic/gin"

	"minecraft-server-statuser/internal/router"
)

func main() {
	engine := gin.Default()

	router.WithRouter(engine)

	engine.Run(":25566")
}
