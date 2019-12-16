package main

import (
"github.com/gin-gonic/gin"
)

func init() {
    LogConf()
}

func test(c *gin.Context) {
    Logger.Info("aaaaaaaa")   //用于测试
}

func main() {
    router := gin.Default()
    router.POST("/test",test)
    router.Run(":8888")
}