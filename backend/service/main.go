package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/tasks", ListTasks)
	e.POST("/tasks", CreateTask)
	e.PUT("/tasks/:id", UpdateTask)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.Logger.Fatal(e.Start(":1323"))
}
