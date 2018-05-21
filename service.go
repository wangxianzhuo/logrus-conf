package logconf

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// Server admin server
type Server int

// Start server start process
//
// POST /log/level exchange level
// 		query param: level
// GET  /log/configs print all log configs as json
func (s Server) Start(address string) {
	e := echo.New()
	e.POST("/log/level", func(c echo.Context) error {
		*Level = c.QueryParam("level")
		LogLevel(*Level)
		return c.JSON(http.StatusOK, "{}")
	})

	e.GET("/log/configs", func(c echo.Context) error {
		buf := new(bytes.Buffer)
		PrintConfigs(buf)
		d, err := ioutil.ReadAll(buf)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, string(d))
	})

	e.HideBanner = true
	e.Start(address)
}
