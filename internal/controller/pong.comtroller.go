package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "maudv")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping .... Pong " + name,
		"uid ":    uid,
	})
}
