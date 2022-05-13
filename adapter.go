package go_socket_io_redis_adapter

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Adapter struct {
	pubConn *redis.PubSubConn
	subConn *redis.PubSubConn

	opts *Options
}

// NewAdapter create new redis adapter. This method used for client API in main function.
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

	a.subConn = &redis.PubSubConn{Conn: conn}
	a.pubConn = &redis.PubSubConn{Conn: conn}

	return a, conn.Close()
}

// NewBroadcast create broadcast for inner server API.
func (a *Adapter) NewBroadcast(nsp string) (*redisBroadcast, error) {
	err := a.subConn.PSubscribe(fmt.Sprintf("%s#%s#*", a.opts.Prefix, nsp))
	if err != nil {
		return nil, err
	}

	uid := newV4UUID()

	rbc := &redisBroadcast{
		rooms:      make(map[string]map[string]Conn),
		requests:   make(map[string]interface{}),
		sub:        a.subConn,
		pub:        a.pubConn,
		key:        fmt.Sprintf("%s#%s#%s", a.opts.Prefix, nsp, uid),
		reqChannel: fmt.Sprintf("%s-request#%s", a.opts.Prefix, nsp),
		resChannel: fmt.Sprintf("%s-response#%s", a.opts.Prefix, nsp),
		nsp:        nsp,
		uid:        uid,
	}

	err = a.subConn.Subscribe(rbc.reqChannel, rbc.resChannel)
	if err != nil {
		return nil, err
	}

	go rbc.dispatch()

	return rbc, nil
}

func (a *Adapter) GetName() string {
	return "redis-adapter"
}
