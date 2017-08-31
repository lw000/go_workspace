package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	auth "melody_test/auth"
	"melody_test/melody"

	"github.com/gin-gonic/gin"
	log "github.com/thinkboy/log4go"
)

func main() {
	log.Debug("version: \"1.0.0\"\n")

	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.POST("/register", func(c *gin.Context) {
		name := c.PostForm("name")
		if len(name) <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": -1, "reason": "name empty"})
			return
		}

		psd := c.PostForm("psd")
		if len(psd) <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": -2, "reason": "psd empty"})
			return
		}

		sign := c.PostForm("sign")
		if len(sign) <= 0 {
			c.JSON(http.StatusOK, gin.H{"code": -3, "reason": "sign empty"})
			return
		}

		uid := rand.Intn(10000) + 10000

		c.JSON(http.StatusOK, gin.H{"code": 0, "uid": uid, "reason": "success"})
	})

	r.GET("/login", func(c *gin.Context) {
		name := c.Query("name")
		if len(name) <= 0 {
			c.Writer.WriteString("name empty")
			return
		}

		psd := c.Query("psd")
		if len(psd) <= 0 {
			c.Writer.WriteString("psd empty")
			return
		}

		c.JSON(http.StatusOK, gin.H{"name": name, "psd": psd})
	})

	r.GET("/logout", func(c *gin.Context) {
		uid := c.Query("uid")
		if len(uid) <= 0 {
			c.Writer.WriteString("uid empty")
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "uid": uid})
	})

	r.GET("/crypto", func(c *gin.Context) {
		data := c.Query("data")
		if len(data) <= 0 {
			c.Writer.WriteString("data empty")
			return
		}
		m := make(map[string]string)
		m["b64"] = auth.Base64Encode([]byte(data))
		m["md5"] = auth.Md5Encode([]byte(data))

		if str, ok := json.Marshal(m); ok == nil {
			c.Writer.WriteString(string(str))
		} else {
			c.Writer.WriteString("0")
		}
	})

	r.POST("/broadcast", func(c *gin.Context) {
		rid := c.PostForm("rid")
		if len(rid) <= 0 {
			c.Writer.WriteString("rid empty")
			return
		}

		uid := c.PostForm("uid")
		if len(uid) <= 0 {
			c.Writer.WriteString("uid empty")
			return
		}

		msg := c.PostForm("msg")
		if len(msg) <= 0 {
			c.Writer.WriteString("msg empty")
			return
		}

		m.BroadcastAll([]byte(msg))

		c.JSON(http.StatusOK, gin.H{"code": 0})
	})

	r.GET("/get_onlines", func(c *gin.Context) {
		rid := c.Query("rid")
		if len(rid) <= 0 {
			c.Writer.WriteString("rid empty")
			return
		}
		uid := c.Query("uid")
		if len(uid) <= 0 {
			c.Writer.WriteString("uid empty")
			return
		}

		var err error
		var count int = 50
		count1 := c.Query("count")
		if len(uid) <= 0 {
			c.Writer.WriteString("count empty")
			return
		} else {
			count, err = strconv.Atoi(count1)
			if err != nil {
				c.Writer.WriteString(err.Error())
				return
			}
		}

		onlines := m.GetOnlines(rid, count)
		if onlines_json, err := json.Marshal(onlines); err == nil {
			c.Writer.Write(onlines_json)
		} else {
			c.Writer.WriteString("error")
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastAll(msg)
	})

	m.HandleConnect(func(h *melody.Hub, s *melody.Session) {
		/*count := */ h.Join(s.Uid, s.Extra)

	})

	m.HandleDisconnect(func(h *melody.Hub, s *melody.Session) {
		/*count := */ h.Leave(s.Uid, s.Extra)
	})

	m.HandleError(func(s *melody.Session, err error) {
		log.Debug("HandleError (rid:%s,uid:%s), error(%v)", s.Rid, s.Uid, err)
	})

	man := &Student{Person{"liwei", 10, "深圳市南山区"}, "深圳大学", 100.00}
	man.test()

	man1 := new(Student)
	man1.Name = "liwei"
	man1.Age = 10
	man1.Address = "深圳市南山区"
	man1.Loan = 100.00
	man1.School = "深圳大学"
	man1.test()

	r.Run(":5000")
}
