package main

import (
	// native
	"net/http"
	"os"

	// external
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

	// yichips
	"yichips/bootstrap"
	customMiddleware "yichips/middleware"
	"yichips/utils/rest"

	// services
	admin "yichips/services/admin/connector"
	authcms "yichips/services/authcms/connector"
	cart "yichips/services/cart/connector"
	product "yichips/services/product/connector"
	transaction "yichips/services/transaction/connector"
	user "yichips/services/user/connector"
)

var DB = bootstrap.DB

/**
 * Main server function
 */
func main() {
	e := echo.New()

	// custom middleware setup
	e.Use(customMiddleware.RequestValidation)

	// logger middleware setup
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS middleware setup
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.HEAD},
	}))

	// pre middleware setup
	e.Pre(middleware.AddTrailingSlash())

	// request validator setup
	e.Validator = &rest.Validator{Validator: validator.New()}

	// service connect by rest handlers
	authcms.Connect(e, "rest")
	admin.Connect(e, "rest")
	product.Connect(e, "rest")
	user.Connect(e, "rest")
	cart.Connect(e, "rest")
	transaction.Connect(e, "rest")

	// server start
	e.GET("/healthz/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Seems Allright !")
	})
	e.Start(":" + os.Getenv("PORT"))

	DB.Close()
}
