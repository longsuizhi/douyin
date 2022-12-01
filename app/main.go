package main

import (
	"github.com/longsuizhi/douyin/server"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	initRouter(c)
	c.Run("127.0.0.1:8899")
}
