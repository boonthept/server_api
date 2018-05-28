//simple_api.go
package main

import (
"net/http"
	"log"
  "fmt"
"github.com/labstack/echo"
)

func main() {
  e := echo.New()
  e.GET("/", func(c echo.Context) error{
    _, err := http.Get("http://localhost:8888")
    if err != nil {
  		log.Fatal(err)
  	}
    fmt.Println("got call from web, pass call to client")
    return c.String(http.StatusOK, "Hello, world")
  })
e.Logger.Fatal(e.Start(":1323"))
}
