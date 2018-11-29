package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main()  {
	e := echo.New()
e.Use(middleware.Recover())
	admin := e.Group("/admin", middleware.BasicAuth(func(s string, s2 string, context echo.Context) (error, bool) {
		return nil, false
	}))
	admin.GET("/userList", func(context echo.Context) error {
		return context.String(200, "hello,world")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
