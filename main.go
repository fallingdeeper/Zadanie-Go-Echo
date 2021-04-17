package main

import (
    "net/http"
    singl "echoapp/singleton"
    "github.com/labstack/echo/v4"
    "fmt"
    table "echoapp/model"
) 

func main() {
    e := echo.New()
    e.POST("/table", func(c echo.Context) error {    

    var object =  singl.getInstance()
    object.OpenS()

    var query = fmt.Sprintf( "INSERT INTO table (Id, Name, Surname) values ('%s', '%s', '%s')", table.Id, table.Name, table.Surname)
    

    singl.CloseS()


    return  c.JSON(table)

   })
   e.Logger.Fatal(e.Start(":1323"))


}
