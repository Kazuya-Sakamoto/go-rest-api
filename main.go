package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	diaryRepository := repository.NewDiaryRepository(db)
	diaryCommentRepository := repository.NewDiaryCommentRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	diaryUsecase := usecase.NewDiaryUsecase(diaryRepository)
	diaryCommentUsecase := usecase.NewDiaryCommentRepository(diaryCommentRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	diaryController := controller.NewDiaryController(diaryUsecase)
	diaryCommentController := controller.NewDiaryCommentController(diaryCommentUsecase)

	e := router.NewRouter(userController, taskController, diaryController, diaryCommentController)
	e.Logger.Fatal(e.Start(":8080"))
}
