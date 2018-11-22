package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/y-ogura/echo-google-cloud-storage/gstorage"
)

// Controller controller
type Controller struct {
	Storage gstorage.GoogleStorage
}

// NewController mount controller
func NewController(e *echo.Echo, storage gstorage.GoogleStorage) {
	handler := &Controller{
		Storage: storage,
	}

	e.GET("/", handler.InputForm)
	e.POST("/upload", handler.Upload)
}

// InputForm input form
func (c *Controller) InputForm(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, formHTML)
}

// Upload upload file
func (c *Controller) Upload(ctx echo.Context) error {
	res, err := c.Storage.Upload(ctx.Request())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

const formHTML = `<!DOCTYPE html>
<html>
  <head>
    <title>Storage</title>
    <meta charset="utf-8">
  </head>
  <body>
    <form method="POST" action="/upload" enctype="multipart/form-data">
      <input type="file" name="file">
      <input type="submit">
    </form>
  </body>
</html>`
