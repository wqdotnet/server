package web

import (
	"fmt"
	"server/tool"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

// Start gin web interface
func Start(Port int32) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/map", func(c *gin.Context) {

		tool.HandleImage(c.Writer, c.Request)
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
	})
	//http.ResponseWriter, reqA *http.Request

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	ginpprof.Wrap(router)

	// ginpprof also plays well with *gin.RouterGroup
	// group := router.Group("/debug/pprof")
	// ginpprof.WrapGroup(group)
	//http://127.0.0.1:8080/debug/pprof/

	router.Run(fmt.Sprintf(":%v", Port))
}
