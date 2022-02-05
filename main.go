package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sultanaliev-s/todo-api/models"
	"github.com/sultanaliev-s/todo-api/user"
)

func GetTasks() []*models.Task {
	return []*models.Task{
		{ID: 1, Description: "Add a database", Author: 1,
			Deadline: time.Now().String(), IsDone: false},
		{ID: 2, Description: "Add tests", Author: 1,
			Deadline: time.Now().String(), IsDone: false},
	}
}

func GetTasksList(c echo.Context) error {
	lt := GetTasks()
	c.JSON(http.StatusOK, lt)
	return nil
}

func main() {
	logger := log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime|log.Llongfile)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/tasks", GetTasksList)

	ur := user.NewRepo()
	us := user.NewService(&ur)
	user.RegisterHandlers(e, &us, logger)

	e.Logger.Fatal(e.Start(":8080"))
}
