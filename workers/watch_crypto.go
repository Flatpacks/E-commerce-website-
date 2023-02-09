package workers

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/jerryan999/CryptoAlert/utils"
	"gith