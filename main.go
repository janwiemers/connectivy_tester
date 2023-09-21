package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// get port
	port := flag.String("port", "8080", "Defines the port the server should run on, default is 8080")
	flag.Parse()

	r := gin.Default()
	r.GET("/perform-request", func(c *gin.Context) {
		url := c.Query("url")
		proto := c.Query("protocol")

		if url == "" || proto == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		body, err := requestFactory(proto, url)
		if err != nil {
			log.Printf("error performing request: %s\t%s\t%s", proto, url, body)
		}
	})
	r.Run(fmt.Sprintf("0.0.0.0:%s", *port))
}

func requestFactory(protocol string, url string) ([]byte, error) {
	log.Printf("performing request: %s\t%s", protocol, url)
	req, err := http.NewRequest(protocol, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("response code: %s", res.Status)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
