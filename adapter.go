package go_socket_io_redis_adapter

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"

	"github.com/gomodule/redigo/redis"
)

type Adapter struct {
	opts *Options
}

// NewAdapter create adapter
func NewAdapter(opts ...OptionsFunc) (*Adapter, error) {
	a := &Adapter{
		opts: NewOptions(),
	}

	for _, o := range opts {
		o(a.opts)
	}

	conn, err := redis.Dial(a.opts.Network, a.opts.Addr)
	if err != nil {
		return nil, err
	}

	return a, conn.Close()
}

func (a *Adapter) NewBroadcast(nsp string) *redisBroadcast {
	addr := a.opts.Addr
	pub, err := redis.Dial(a.opts.Network, addr)
	if err != nil {
		//todo write to log message
		return nil, err
	}

	sub, err := redis.Dial(a.opts.Network, addr)
	if err != nil {
		//todo write to log message
		return nil, err
	}

	subConn := &redis.PubSubConn{Conn: sub}
	pubConn := &redis.PubSubConn{Conn: pub}

	if err = subConn.PSubscribe(fmt.Sprintf("%s#%s#*", a.opts.Prefix, nsp)); err != nil {
		//todo write to log message
		return nil, err
	}

	uid := newV4UUID()
	rbc := &redisBroadcast{
		rooms:      make(map[string]map[string]socketio.Conn),
		requests:   make(map[string]interface{}),
		sub:        subConn,
		pub:        pubConn,
		key:        fmt.Sprintf("%s#%s#%s", a.opts.Prefix, nsp, uid),
		reqChannel: fmt.Sprintf("%s-request#%s", a.opts.Prefix, nsp),
		resChannel: fmt.Sprintf("%s-response#%s", a.opts.Prefix, nsp),
		nsp:        nsp,
		uid:        uid,
	}

	if err = subConn.Subscribe(rbc.reqChannel, rbc.resChannel); err != nil {
		//todo write to log message
		return nil, err
	}

	go rbc.dispatch()

	return rbc
}

func (a *Adapter) GetName() string {
	return "redis-adapter"
}
