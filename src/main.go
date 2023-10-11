//package main
//
//import (
//	"github.com/gin-gonic/gin"
//
//	"net/http"
//)
//
//func main() {
//	rs := gin.Default()
//	//r := gin.New()
//	rs.Use(gin.Logger())
//	rs.Use(gin.Recovery())
//
//	rs.GET("/ping", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": "算法",
//		})
//	})
//	rs.Run(":8000")
//}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "反击!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
