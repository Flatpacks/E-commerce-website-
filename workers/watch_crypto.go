package workers

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/jerryan999/CryptoAlert/utils"
	"github.com/labstack/gommon/log"
)

var QUEUE_KEY_WATCH_CRYPTO = "tasks.send_email"
var QUEUE_ADDR_WATCH_CRYPTO = "localhost:6379"

// redis
var rdb *redis.Client

func WatchCryptoWorker() {
	ctx := context.TODO()

	// we only care about crypto here
	crypto := "bitcoin"

	conn, _, err := websocket.DefaultDialer.Dial("wss://ws.coincap.io/prices?assets="+crypto, nil)
	if err != nil {
		log