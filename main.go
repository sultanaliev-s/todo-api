package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID          uint64 `json:"id"`
	Description string `json:"description"`
	Author      uint64 `json:"author"`
	Deadline    string `json:"deadline"`
	IsDone      bool   `json:"isDone"`
}

func GetTasks() []*Task {
	return []*Task{
		{1, "Add a database", 1, time.Now().String(), false},
		{2, "Add tests", 1, time.Now().String(), false},
	}
}

func GetTasksList(c echo.Context) error {
	lt := GetTasks()
	c.JSON(http.StatusOK, lt)
	return nil
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/tasks", GetTasksList)

	e.Logger.Fatal(e.Start(":8080"))
}
