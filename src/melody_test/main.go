package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	log "github.com/thinkboy/log4go"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/login", func(c *gin.Context) {
		name := c.Query("name")
		psd := c.Query("psd")
		c.JSON(http.StatusOK, gin.H{"name": name, "psd": psd})
	})

	r.POST("/broadcast", func(c *gin.Context) {
		rid := c.PostForm("rid")
		uid := c.PostForm("uid")
		msg := c.PostForm("msg")
		m.Broadcast([]byte(msg))
		c.JSON(http.StatusOK, gin.H{"rid": rid, "uid": uid, "msg": msg})
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return true
		})
	})

	log.Debug("start")

	r.Run(":5000")
}
