package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":"OK",
			"message": "登陆成功",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}



