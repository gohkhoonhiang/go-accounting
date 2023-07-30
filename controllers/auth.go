package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-accounting/models"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "id"

type LoginUserParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Reference: https://gowebexamples.com/password-hashing/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CurrentUser(c *gin.Context) any {
	user, _ := c.Get(identityKey)
	return user.(*models.User).Username
}

func ValidateUser(c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	return len(claims[identityKey].(string)) > 0
}

// Reference: https://github.com/Depado/gin-auth-example/blob/master/main.go
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		valid := ValidateUser(c)

		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
			return
		}

		c.Next()
	}
}

func InitAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("SESSION_SECRET")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: claims[identityKey].(string),
			}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var input LoginUserParam
			if err := c.ShouldBindJSON(&input); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			var user models.User
			if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
				return "", jwt.ErrFailedAuthentication
			}

			if match := CheckPasswordHash(input.Password, user.Password); !match {
				return "", jwt.ErrFailedAuthentication
			}

			return &user, nil
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.User); ok {
				return true
			}

			return false
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			log.Println(message)
			c.JSON(code, gin.H{
				"error": "Invalid credentials",
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	return authMiddleware, err
}
