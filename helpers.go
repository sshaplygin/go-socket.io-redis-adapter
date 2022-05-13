package go_socketio_redis_adapter

import "github.com/gofrs/uuid"

func newV4UUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
