package service

import (
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
	"github.com/jerryan999/CryptoAlert/model"
	"github.com/jerryan999/CryptoAlert/utils"
	"github.com/labstack/gommon/log"
)

type AlertService struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAlertService(db *sql.DB, rdb *redis.Client) *AlertService {
	return &AlertService{db: db, rdb: rdb}
}

func (s *AlertService) GetAlertByID(id int64) (*model.Alert, error) {
	return model.GetAlertByID(s.db, id)
}

func (s *AlertService) AddAlert(alert *model.Alert) (*model.Alert, error) {
	// Add alert to database
	alert, err := model.SaveAlert(s.db, alert)
	if err != nil {
		return alert, errors.New("saving alert to DB Failed")
	}
	log.Info("Alert saved to DB")

	// save alert to redis sorted set
	key := utils.GetAlertQueueKey(alert.Crypto, alert