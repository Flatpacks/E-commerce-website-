
package server

import (
	"database/sql"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jerryan999/CryptoAlert/database"

	"github.com/jerryan999/CryptoAlert/model"
	"github.com/jerryan999/CryptoAlert/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var (
	// persistent the alerts in a sqlite database
	db *sql.DB = database.Initialize("./database/db.sqlite")

	// store the alerts in a redis database for further processing
	rdb *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})
)

type AlertController struct {
	alertService *service.AlertService
}

func NewAlertController() *AlertController {
	return &AlertController{
		alertService: service.NewAlertService(db, rdb),
	}
}

func (ctl *AlertController) AddAlert(c echo.Context) error {
	data := GetJSON(c)
	alert := &model.Alert{
		Crypto:    data["crypto"].(string),
		Direction: data["direction"].(bool), // true for above, false for below
		Price:     data["price"].(float64),
	}
	if alert.Price < 0 || alert.Crypto != "bitcoin" {
		log.Warn("Invalid data", alert)
		return c.JSON(http.StatusOK, NewInvalidDataErrorResponse())
	}
	alert, err := ctl.alertService.AddAlert(alert)
	if err != nil {
		return c.JSON(http.StatusOK, NewParseRequestErrorResponse())
	}
	return c.JSON(http.StatusOK, NewSuccessResponse(alert))