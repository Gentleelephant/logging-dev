package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

type Message struct {
}

func main() {

	r := gin.Default()

	r.POST("/log/http", receiveLogs)

	r.Run(":9999")

	ch := make(chan bool)

	<-ch
}

func receiveLogs(c *gin.Context) {

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%v\n", string(data))

	// TODO : 接收vector发来的日志

	// TODO : 反序列化，但是使用什么结构体来接收？

	// TODO : 将日志写入对应channel

}
