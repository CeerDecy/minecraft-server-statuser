package handle

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"minecraft-server-statuser/internal/ping-server"
)

func StatusHandle(ctx *gin.Context) {

	var (
		host        = "127.0.0.1"
		port uint16 = 25565
		opts        = &ping_server.StatusOptions{
			Query:   true,
			Timeout: time.Second * 5,
		}
	)

	response, expiresAt, err := ping_server.GetJavaStatus(host, port, opts)

	defer func() {
		if err != nil {
			logrus.Error(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "can't get server status",
			})
		}
	}()

	if err != nil {
		return
	}

	ctx.Set("X-Cache-Hit", strconv.FormatBool(expiresAt != 0))

	if expiresAt != 0 {
		ctx.Set("X-Cache-Time-Remaining", strconv.Itoa(int(expiresAt.Seconds())))
	}
	ctx.JSON(200, response)
}
