package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UserInfo message
type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func indexHandler(c *gin.Context) {
	// time.Sleep(1 * time.Second) 
	fmt.Println("indexHandler")
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func m1(c *gin.Context) {
	start := time.Now()
	fmt.Println("m1 in ...")
	c.Next()  //调用后续的处理函数
	c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	fmt.Println("m1 out ...")
}
func main() {
	r := gin.Default()
	r.Use(m1) // 全局注册中间件函数m1
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "shop"}) })
	r.Run(":9090")
}
