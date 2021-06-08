package go_socket_io_redis_adapter

// Logger logs messages with different levels.
type Logger interface {
	Errorf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}
