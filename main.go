package main

import (
    "net/http"
    "echoapp/singleton"
    "github.com/labstack/echo/v4"
    "echoapp/model"
) 

func main() {
    e := echo.New()
    e.GET("/table", func(c echo.Context) error {    

    var object =  singleton.GetInstance()
    object.OpenS()

    var table = model.Table{
	    Id:	"1",
	    Name:	"Olga",
	    Surname:	"Petryk",
    }

    object.CloseS()


    return  c.JSON(http.StatusOK, table)

   })
   e.Logger.Fatal(e.Start(":1323"))

}
