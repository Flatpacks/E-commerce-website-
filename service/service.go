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

func NewAlertService(db *sql.DB, rdb *redi