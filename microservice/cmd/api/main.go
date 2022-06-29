package main

import (
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/user"
	"course/internal/user/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"message": "hello world",
		})
	})

	db := database.NewDatabaseConn()
	exerciseService := exercise.NewExerciseService(db)
	// userDbRepo := repository.NewDBRepository(db)
	mcsrvRepo := repository.NewMicroserviceRepo()
	userService := user.NewUserService(mcsrvRepo)
	// exercises
	route.GET("/exercises/:id", middleware.Authentication(userService), exerciseService.GetExercise)
	route.GET("/exercises/:id/score", middleware.Authentication(userService), exerciseService.GetUserScore)

	route.Run(":8000")
}
