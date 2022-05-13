package go_socketio_redis_adapter

// Options is configuration to create new adapter.
type Options struct {
	Addr    string
	Prefix  string
	Network string //notice: redis supported tcp|unix network
}

// newOptions create default options.
func newOptions() *Options {
	return &Options{
		Addr:    "127.0.0.1:6379",
		Prefix:  "socket.io",
		Network: "tcp",
	}
}

// OptionsFunc as option interface.
type OptionsFunc func(o *Options)

// WithAddrOptions set custom connect addr.
func WithAddrOptions(addr string) OptionsFunc {
	return func(o *Options) {
		o.Addr = addr
	}
}

// WithPrefixOptions set custom prefix for redis key.
func WithPrefixOptions(prefix string) OptionsFunc {
	return func(o *Options) {
		o.Prefix = prefix
	}
}

// WithNetworkOptions set custom connection network type.
func WithNetworkOptions(network string) OptionsFunc {
	return func(o *Options) {
		o.Network = network
	}
}
