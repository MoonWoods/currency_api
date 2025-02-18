package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MoonWoods/currency_api/pkg/currency"
)

var server *gin.Engine
var port = 8080

func SetupRouter() {
	server = gin.Default()

	server.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	server.GET("/all", func(c *gin.Context) {
		data, err := currency.GetAll()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, data)
	})

	server.GET("/exchange/:currency", func(c *gin.Context) {
		curr := c.Param("currency")
		data, err := currency.GetExchanges(curr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, data)
	})
}

func StartRouter() {
	slog.Info(fmt.Sprintf("Listening on port '%v'", port))
	server.Run(fmt.Sprintf(":%v", port))
}
