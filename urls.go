package main

import (
 "mymodule/server"
 "mymodule/boards"
 "github.com/labstack/echo/v4"
 "net/http"
 "mymodule/database"
 "html/template" // Import the template package
)

func HealthCheck(c echo.Context) error {
 return c.String(http.StatusOK, "OK")
}

func GetBoard(c echo.Context) error {
 db := database.GetDB() 
 defer db.Close()
 tmpl := template.Must(template.ParseFiles("templates/boards.html"))
 boards, err := boards.GetAllBoards(db, tmpl) // Pass db and tmpl to GetAllBoards
 if err != nil {
  return c.String(http.StatusInternalServerError, err.Error())
 }
 return c.JSON(http.StatusOK, boards)
}

func RegisterRoutes(e *echo.Echo) {
 e.GET("/health", server.HealthCheck)
 e.GET("/boards", GetBoard) // Use GetBoard from urls.go
}

func main() {
 database.InitializeDB()
 e := echo.New()
 RegisterRoutes(e)
 e.Logger.Fatal(e.Start(":8080"))
}