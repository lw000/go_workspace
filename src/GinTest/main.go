// GinTest project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method //请求方法
		origin := c.Request.Header.Get("Origin")
		//请求头部
		var headerKeys []string
		// 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}

		headerStr := strings.Join(headerKeys, ", ")
		if headerStr == "" {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		} else {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		}

		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                                                                                                                                                                                                                                                     // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")                                                                                                                                                                                                                               //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求 // header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma") // 允许跨域设置 可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")                                                                                                           // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                                                                                                                                     // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                                                                                                                            // 跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                                                                                                                                        // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		// 处理请求
		c.Next()
		// 处理请求
	}
}

func main() {
	r := gin.Default()
	r.Use(Cors())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world.")
	})

	r.GET("/add", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")
		fmt.Printf("a=%s, b=%s\n", a, b)
		ia, err := strconv.Atoi(a)
		if err != nil {

		}

		ib, err := strconv.Atoi(b)
		if err != nil {

		}

		c.JSON(http.StatusOK, gin.H{"data": ia + ib})
	})

	r.GET("/test", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		//		c.String(http.StatusOK, "Hello [%s %s]", firstname, lastname)
		c.JSON(http.StatusOK, gin.H{"firstname": firstname, "lastname": lastname})
	})

	r.POST("/error", func(c *gin.Context) {
		c.String(http.StatusNotFound, "and error hapenned.")
	})

	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		filename := header.Filename

		fmt.Println(header.Filename)

		out, err := os.Create("./tmp/" + filename + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/post", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})

		v1.POST("/post1", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/post", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})

		v2.POST("/post1", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})
	}

	r.Run(":8888")
}
