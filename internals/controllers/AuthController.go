package controllers

import (
	"net/http"

	"github.com/ArifulProtik/socialist/internals/views/pages"
	"github.com/labstack/echo/v4"
)

func (c *Controller) Signup(ctx echo.Context) error {
	return c.Render(ctx, http.StatusOK, pages.Index())
}
