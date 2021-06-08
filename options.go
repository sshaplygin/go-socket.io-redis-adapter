package go_socket_io_redis_adapter

// Options is configuration to create new adapter
type Options struct {
	Addr    string
	Prefix  string
	Network string
}

func NewOptions() *Options {
	return &Options{
		Addr:    "127.0.0.1:6379",
		Prefix:  "socket.io",
		Network: "tcp",
	}
}

type OptionsFunc func(o *Options)

// WithAddrOptions
func WithAddrOptions(addr string) OptionsFunc {
	return func(o *Options) {
		o.Addr = addr
	}
}

// WithPrefixOptions
func WithPrefixOptions(prefix string) OptionsFunc {
	return func(o *Options) {
		o.Prefix = prefix
	}
}

// WithNetworkOptions
func WithNetworkOptions(network string) OptionsFunc {
	return func(o *Options) {
		o.Network = network
	}
}
