package main

import (
	"github.com/k-ueki/class-reputation/src/server/db"
	"github.com/labstack/echo"
)

func main() {
	db := db.NewDBStruct().Connect()
	defer db.CLose()
	fmt.Println("db,", db)

	e := echo.New()

	testController
	e.GET("/hc", testController.Test())

	e.Start(":9000")
}
