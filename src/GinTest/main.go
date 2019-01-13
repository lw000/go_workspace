// GinTest project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	//	"github.com/Salvatore-Giordano/gin-redis-ip-limiter"
	"github.com/gin-gonic/gin"
	//	"github.com/utrack/gin-csrf"

	"GinTest/middleware"

	"github.com/muesli/cache2go"
)

func TestCache() {
	c := cache2go.Cache("levi")
	c.Add("levi", time.Second*time.Duration(5), "1111111111111111111")
	c.Add("levi1", time.Second*time.Duration(5), "1111111111111111111")
	c.Add("levi2", time.Second*time.Duration(5), "1111111111111111111")
	c.Add("levi3", time.Second*time.Duration(5), "1111111111111111111")
	c.Add("levi4", time.Second*time.Duration(5), "1111111111111111111")
	c.Add("levi5", time.Second*time.Duration(5), "1111111111111111111")

	c.Foreach(func(v interface{}, item *cache2go.CacheItem) {
		fmt.Printf("%+v\n", item.Data().(string))
	})
}

func main() {
	TestCache()

	r := gin.Default()
	r.Use(middleware.Cros())
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

	r.Run(":8888")
}
