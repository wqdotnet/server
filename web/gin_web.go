package web

import (
	"fmt"
	"server/gserver/commonstruct"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Start gin web interface
func Start(Port int32) {
	log.Info("Start [Web Http]")
	//禁用控制台颜色，在将日志写入文件时不需要控制台颜色
	//gin.DisableConsoleColor()

	//如果需要控制台输出带有颜色的字体，请使用下面代码
	gin.ForceConsoleColor()

	//如果需要将日志写入文件，请使用以下代码
	//创建日志文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	//如果需要将日志输出到控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(os.Stdout)

	//如果需要同时将日志写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()
	router.Use(logger(), gin.Recovery())

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/refreshCfg", refreshCfg)

	router.GET("/map", func(context *gin.Context) {
		//tools.HandleImage(c.Writer, c.Request)
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if commonstruct.ServerCfg.OpenWS {
		hub := newHub()
		go hub.run()

		router.GET("/ws", func(context *gin.Context) {
			WsClient(hub, context)
		})
	}

	//http.ResponseWriter, reqA *http.Request

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	//ginpprof.Wrap(router)

	// ginpprof also plays well with *gin.RouterGroup
	// group := router.Group("/debug/pprof")
	// ginpprof.WrapGroup(group)
	//http://localhost:8080/debug/pprof/

	router.Run(fmt.Sprintf(":%v", Port))

}

//刷新配置
func refreshCfg(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// 日志中间件
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.Request.Host
		log.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqURI)
	}
}
