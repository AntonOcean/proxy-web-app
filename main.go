package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"time"
)

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func proxyStart(c *gin.Context) {
	r := c.Request
	r.RequestURI = ""

	client := http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusServiceUnavailable, err.Error())
		return
	}
	defer resp.Body.Close()

	copyHeader(c.Writer.Header(), resp.Header)
	c.Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

func heavyBitch() {
	time.Sleep(5 * time.Second)
	println("i am done")
}

func getRequest(c *gin.Context) {
	go heavyBitch()
	fmt.Println("success")
	c.String(http.StatusOK, "hello")
}

func main() {
	proxy := gin.Default()
	proxy.Any("/:any", proxyStart)
	go func() {
		log.Fatal(proxy.Run("0.0.0.0:9090"))
	}()

	proxyAPI := gin.Default()
	proxyAPI.GET("/hello", getRequest)
	log.Fatal(proxyAPI.Run("0.0.0.0:8080"))
}
