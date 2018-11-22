package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/y-ogura/echo-google-cloud-storage/controller"
	"github.com/y-ogura/echo-google-cloud-storage/gstorage"
)

func main() {
	e := echo.New()

	storage := gstorage.NewGoogleStorage()

	controller.NewController(e, storage)

	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
