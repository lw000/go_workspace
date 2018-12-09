package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"melody_test/auth"
	"net/http"
	"os"
	"strconv"

	"melody_test/melody"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"

	//	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	log "github.com/thinkboy/log4go"

	"melody_test/middleware"
)

func main() {
	r := gin.Default()

	//	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//	r.Use(sessions.Sessions("mysession", store))

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Use(middleware.Cros())

	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/add", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")
		ia, err := strconv.Atoi(a)
		if err != nil {

		}

		ib, err := strconv.Atoi(b)
		if err != nil {

		}

		c.JSON(http.StatusOK, gin.H{
			"a":    ia,
			"b":    ib,
			"data": ia + ib,
		})
	})

	r.POST("/sub", func(c *gin.Context) {
		a := c.Request.FormValue("a")
		b := c.Request.FormValue("b")
		ia, err := strconv.Atoi(a)
		if err != nil {

		}

		ib, err := strconv.Atoi(b)
		if err != nil {

		}

		c.JSON(http.StatusOK, gin.H{
			"a":    ia,
			"b":    ib,
			"data": ia - ib,
		})

	})

	r.GET("/test", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		//		c.String(http.StatusOK, "Hello [%s %s]", firstname, lastname)
		c.JSON(http.StatusOK, gin.H{"firstname": firstname, "lastname": lastname})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename

		fmt.Println(header.Filename)

		out, err := os.Create("./tmp/" + filename + ".png")
		if err != nil {
			log.Error(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Error(err)
		}
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/post", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"xxx": 1111})
		})

		v1.POST("/post1", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})
	}

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.POST("/register", func(c *gin.Context) {
		name := c.PostForm("name")
		if len(name) == 0 {
			c.JSON(http.StatusOK, gin.H{"code": -1, "reason": "name empty"})
			return
		}

		psd := c.PostForm("psd")
		if len(psd) == 0 {
			c.JSON(http.StatusOK, gin.H{"code": -2, "reason": "psd empty"})
			return
		}

		sign := c.PostForm("sign")
		if len(sign) == 0 {
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
		m["sha1"] = auth.Sha1([]byte(data))
		m["sha224"] = auth.Sha224([]byte(data))
		m["sha256"] = auth.Sha256([]byte(data))
		m["sha512"] = auth.Sha512([]byte(data))

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

		var (
			err    error
			count  int    = 50
			count1 string = ""
		)
		count1 = c.Query("count")
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
		count := h.Join(s.Uid, s.Extra)
		if count > 0 {

		}
	})

	m.HandleDisconnect(func(h *melody.Hub, s *melody.Session) {
		count := h.Leave(s.Uid, s.Extra)
		if count > 0 {

		}
	})

	m.HandleError(func(s *melody.Session, err error) {
		log.Debug("HandleError (rid:%s,uid:%s), error(%v)", s.Rid, s.Uid, err)
	})

	man := &Student{Person{"liwei", 10, "深圳市南山区"}, "深圳大学", 100.00}
	man.test()

	r.Run(":5000")
}
