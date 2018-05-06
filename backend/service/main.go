package main

import (
	"net/http"
	"strconv"

	"github.com/anraku/TodoList/backend/database"
	"github.com/anraku/TodoList/backend/helper"
	"github.com/anraku/TodoList/backend/model"
	"github.com/anraku/TodoList/backend/model/tasks"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/tasks", listTasks)
	e.POST("/tasks", createTask)
	e.PUT("/tasks/:id", updateTask)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.Logger.Fatal(e.Start(":1323"))
}

func listTasks(c echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()
	tasks, err := tasks.New(db).FindAll()
	if err != nil {
		return err
	}
	tasksResponse := make([]helper.ResponseMap, len(tasks))
	for i, task := range tasks {
		tasksResponse[i] = task.ToResponseMap()
	}
	return helper.JSONResponseArray(c, http.StatusOK, tasksResponse)
}

func createTask(c echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	c.Request().ParseForm()
	text := c.Request().Form["text"][0]
	done, err := strconv.ParseBool(c.Request().Form["done"][0])
	if err != nil {
		return err
	}
	m := model.Task{
		Text: text,
		Done: done,
	}
	task, err := tasks.New(db).Create(m)
	if err != nil {
		return err
	}
	return helper.JSONResponseObject(c, http.StatusOK, &task)
}

func updateTask(c echo.Context) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	id := c.Param("id")
	c.Request().ParseForm()
	done, err := strconv.ParseBool(c.Request().Form["done"][0])
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	m := model.Task{
		ID:   i,
		Done: done,
	}
	task, err := tasks.New(db).Update(m)
	if err != nil {
		return err
	}

	return helper.JSONResponseObject(c, http.StatusOK, &task)
}
