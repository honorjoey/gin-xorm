package ginmiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/honorjoey/gin-xorm/log"
	"net/http"
	"time"
)

type MLog struct {
	log.Log
}

var mLog MLog

func init() {
	mLog.Tag = "middleware"
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}

// NoRoute is a middleware to handle page not found during requests
func NoRoute(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
}

func RequestTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	mLog.Info("request time", since)
}
