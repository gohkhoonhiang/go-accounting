package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-accounting/controllers"
	"github.com/go-accounting/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Use(static.Serve("/", static.LocalFile("./assets", true)))
	r.Use(static.Serve("/favicon.ico", static.LocalFile("./assets/favicon.ico", true)))

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/budgets", controllers.FindBudgets)
	r.POST("/budgets", controllers.CreateBudget)
	r.GET("/budgets/:id", controllers.FindBudget)
	r.PATCH("/budgets/:id", controllers.UpdateBudget)
	r.DELETE("/budgets/:id", controllers.DeleteBudget)

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "9090"
	}
	r.Run(fmt.Sprintf(":%v", port))
}
