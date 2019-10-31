package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
)

//func main() {
//  e := echo.New()
//  e.GET("/", func(c echo.Context) error {
//    return c.String(http.StatusOK, "Hello, World!")
//  })
//  e.Logger.Fatal(e.Start(":1323"))
//}

func main() {
  e := echo.New()
  
  // logging
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
  track := func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
      println("request to /users")
      return next(c)
    }
  }
  e.GET("/users", func(c echo.Context) error {
    return c.String(http.StatusOK, "/users")
  }, track)
  
  // basic auth
  g := e.Group("/admin")
  g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
    if username == "hawksnowlog" && password == "secret" {
      return true, nil
    }
    return false, nil
  }))
  g.GET("/users", func(c echo.Context) error {
    return c.String(http.StatusOK, "/admin/users")
  }, track)
  
  e.Logger.Fatal(e.Start(":1323"))
  
}