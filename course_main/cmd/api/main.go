package main

import (
	"course/internal/answer"
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/question"
	"course/internal/user"

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
	answerService := answer.NewAnswerService(db)
	questionService := question.NewQuestionService(db)
	userService := user.NewUserService(db)
	// exercises
	route.GET("/exercises/:id", middleware.Authentication(userService), exerciseService.GetExercise)
	route.GET("/exercises/:id/score", middleware.Authentication(userService), exerciseService.GetUserScore)
	route.POST("/exercises/:id/questions", middleware.Authentication(userService), exerciseService.CreateQuestions)
	route.POST("/exercises/:id/questions/:questionID/answer", middleware.Authentication(userService), exerciseService.CreateAnswer)

	// answers
	route.GET("/answers/:id", middleware.Authentication(userService), answerService.GetAnswer)

	// questions
	route.GET("/questions/:id", middleware.Authentication(userService), questionService.GetQuestion)

	// user
	route.POST("/register", userService.Register)
	route.POST("/login", userService.Login)
	route.Run(":8000")
}
