package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DataMhs struct {
	ID   int64  `json:"nik"`
	Nama string `json:"nama" validate:"required"`
}

func main() {
	e := echo.New()
	fmt.Println("Hai")

	e.GET("/datamhs", HelloController)
	//e.GET("/datamhs/:id", HelloController)

	e.Start(":8085")
}

func HelloController(c echo.Context) error {
	var data DataMhs
	_ = c.Bind(&data)
	if data.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "ID must exist",
			"data":          data,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
