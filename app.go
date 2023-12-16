package main

import (
	"chinet.work/chilive/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"yajuu": "senpai!",
		})
	})
	
	app.POST("/stream/get", func(c *gin.Context) {
		room, _ := c.GetQuery("room")
		if room != "" {
			resp, err := http.Get(config.LivegoAPIServerURL + "?room=" + room)
			if err != nil {
				panic("Can not contact the LiveGo Server!")
			}
			
		}
	})
	
	app.Run(":" + strconv.Itoa(config.ListenPort))
}
