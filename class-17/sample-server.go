package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	OK          bool   `json:"ok"`
	ServiceName string `json:"serviceName"`
}

func main() {
	r := gin.Default()

	r.GET("/", HealthHanlder())
	r.GET("/home", HomeHandler())
	r.GET("/friends/:id")
	// paht parameter

	// GET /friends?limit=100&offset=1

	fmt.Println("starting the server...")
	r.Run(":8080")
	// localhost:8080
}

func HomeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")

		c.GetQuery()


		c.JSON(http.StatusOK, "Welcome!")
	}
}

func HealthHanlder() gin.HandlerFunc {
	res := &response{
		OK:          true,
		ServiceName: "demo-service",
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &res)
	}
}
