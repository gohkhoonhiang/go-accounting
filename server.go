package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	r.Use(cors.New(config))

	authMiddleware, err := controllers.InitAuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.POST("/signup", controllers.CreateUser)
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/logout", authMiddleware.LogoutHandler)

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())

	api := r.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	api.Use(controllers.RequireAuth())
	{
		api.GET("/current_user", controllers.GetCurrentUser)
		api.PATCH("/current_user", controllers.UpdateUser)

		api.GET("/budgets", controllers.FindBudgets)
		api.POST("/budgets", controllers.CreateBudget)
		api.GET("/budgets/:id", controllers.FindBudget)
		api.PATCH("/budgets/:id", controllers.UpdateBudget)
		api.DELETE("/budgets/:id", controllers.DeleteBudget)
	}

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "9090"
	}
	r.Run(fmt.Sprintf(":%v", port))
}
