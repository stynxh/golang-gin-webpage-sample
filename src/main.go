package main

import (      
    "net/http"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
	router.Static("css", "./css")
	router.Static("font", "./font")
	router.Static("images", "./images")
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
		/*c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
		*/
    })
	router.GET("/about", func(c *gin.Context) {
        c.HTML(http.StatusOK, "about.html", gin.H{})		
    })
	router.GET("/project", func(c *gin.Context) {
        c.HTML(http.StatusOK, "project.html", gin.H{})		
    })

    router.Run(":80")
}
