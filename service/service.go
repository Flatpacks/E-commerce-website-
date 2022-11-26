package service

import (
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
	"github.com/jerryan999/CryptoAlert/model"
	"github.com/jerryan999/CryptoAlert/