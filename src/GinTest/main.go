// GinTest project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world.")
	})

	router.GET("/test", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		//		c.String(http.StatusOK, "Hello [%s %s]", firstname, lastname)
		c.JSON(http.StatusOK, gin.H{"firstname": firstname, "lastname": lastname})
	})

	router.POST("/error", func(c *gin.Context) {
		c.String(http.StatusNotFound, "and error hapenned.")
	})

	router.POST("/upload", func(c *gin.Context) {
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

	v1 := router.Group("/v1")
	{
		v1.POST("/post", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})

		v1.POST("/post1", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})

		v1.POST("/post2", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"xxx": 1111})
		})
	}

	router.Run(":8080")
}
