package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/sse", SSEWithChannel)
	r.GET("/sseA", SSE)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server run failed")
		return
	}

}

// sse demo
func SSE(c *gin.Context) {
	closeNotify := c.Request.Context().Done()

	go func() {
		<-closeNotify
		fmt.Println("SSE Closed")
		return
	}()

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", "Hello, SSE!")

		for i := 0; i < 5; i++ {
			c.SSEvent("message", fmt.Sprintf("SSE! %d", i))
			time.Sleep(500 * time.Millisecond)
		}

		c.SSEvent("message", "Goodbye, SSE!")
		fmt.Println("SSE Closed")

		return false
	})
}

// send message with channel
func SSEWithChannel(c *gin.Context) {
	closeNotify := c.Request.Context().Done()

	messageChan := make(chan string)
	defer close(messageChan)

	go func() {
		<-closeNotify
		fmt.Println("SSE Closed")
		return
	}()

	go SendMessage(messageChan)

	c.Stream(func(w io.Writer) bool {
		select {
		case message := <-messageChan:
			c.SSEvent("message", message)
			return true
		case <-closeNotify:
			fmt.Println("SSE Closed")
			return false
		}
	})

}

// messageChan
func SendMessage(messageChan chan string) {
	for i := 0; i < 5; i++ {
		messageChan <- fmt.Sprintf("SSE! %d", i)
		time.Sleep(500 * time.Millisecond)
	}
}
