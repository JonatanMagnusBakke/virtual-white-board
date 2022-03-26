package main

import (
	"log"
	"virtual-white-board-service/internal/backends/comment"
	"virtual-white-board-service/internal/backends/post"
	"virtual-white-board-service/internal/backends/user"
	"virtual-white-board-service/internal/database"
	"virtual-white-board-service/internal/flags"
	"virtual-white-board-service/internal/httphandlers"
	"virtual-white-board-service/internal/httphandlers/httpcomment"
	"virtual-white-board-service/internal/httphandlers/httppost"
	"virtual-white-board-service/internal/httphandlers/httpuser"
	"virtual-white-board-service/internal/httphandlers/middleware"
	"virtual-white-board-service/internal/services"

	"github.com/gin-gonic/gin"
)

//CORSMiddleware handles cors problem
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	flags, err := flags.ParseFlags()
	if err != nil {
		log.Fatalf("Missing flag: %s", err)
	}

	db, err := database.New(flags)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	err = database.CreateTestUsers(db)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	jwtService := services.JWTAuthService()

	e := gin.New()
	e.SetTrustedProxies(nil)

	if flags.AllowCORS {
		e.Use(CORSMiddleware())
	}
	//No auth needed
	e.POST(httphandlers.LoginRoute, httpuser.NewLoginHandler(user.NewLogin(db, jwtService)))

	e.Use(middleware.AuthMiddleware(jwtService))

	//Auth needed
	e.GET(httphandlers.PostsRoute, httppost.NewListHandler(post.NewLister(db)))

	e.POST(httphandlers.PostsRoute, httppost.NewInsertHandler(post.NewInserter(db)))

	e.POST(httphandlers.CommentRoute, httpcomment.NewInsertHandler(comment.NewInserter(db)))

	err = e.Run(":" + flags.APIPort)
	if err != nil {
		log.Fatal("Could not run server")
	}
}
