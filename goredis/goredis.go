package goredis

import (
	bootflag "github.com/al8n/micro-boot/flag"
	"github.com/al8n/micro-boot/internal/utils"
	// "github.com/go-redis/redis/v8"
	"time"
)

var (
	defaultGoRedisFlagsPrefix = "goredis"
	defaultGoRedisNetwork = "tcp"
)

func SetDefaultGoRedisFlagsPrefix(val string)  {
	defaultGoRedisFlagsPrefix = val
}

func SetDefaultGoRedisNetwork(val string)  {
	defaultGoRedisNetwork = val
}

type Options struct {
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string   `json:"network" yaml:"network" mapstructure:"network"`
	// host:port address.
	Addr string `json:"addr" yaml:"addr" mapstructure:"addr"`

	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `json:"username" yaml:"username" mapstructure:"username"`
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string  `json:"password" yaml:"password" mapstructure:"password"`

	// Database to be selected after connecting to the server.
	DB int `json:"db" yaml:"db" mapstructure:"db"`

	// Maximum number of retries before giving up.
	// Default is 3 retries.
	MaxRetries int `json:"max-retries" yaml:"max-retries" mapstructure:"max-retries"`
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration `json:"min-retry-backoff" yaml:"min-retry-backoff" mapstructure:"min-retry-backoff"`
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration `json:"max-retry-backoff" yaml:"max-retry-backoff" mapstructure:"max-retry-backoff"`

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `json:"dial-timeout" yaml:"dial-timeout" mapstructure:"dial-timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout time.Duration `json:"read-timeout" yaml:"read-timeout" mapstructure:"read-timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout time.Duration `json:"write-timeout" yaml:"write-timeout" mapstructure:"write-timeout"`

	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int `json:"pool-size" yaml:"pool-size" mapstructure:"pool-size"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `json:"min-idle-conns" yaml:"min-idle-conns" mapstructure:"min-idle-conns"`
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge time.Duration `json:"max-conn-age" yaml:"max-conn-age" mapstructure:"max-conn-age"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout time.Duration `json:"pool-timeout" yaml:"pool-timeout" mapstructure:"pool-timeout"`
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout time.Duration `json:"idle-timeout" yaml:"idle-timeout" mapstructure:"idle-timeout"`
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency time.Duration `json:"idle-check-frequency" yaml:"idle-check-frequency" mapstructure:"idle-check-frequency"`

	// Enables read only queries on slave nodes.
	ReadOnly bool `json:"read-only" yaml:"read-only" mapstructure:"read-only"`

	// TLS Config to use. When set TLS will be negotiated.
	// TLSConfig *tls.Config

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-" mapstructure:",omitempty"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-" mapstructure:",omitempty"`
}

func (o *Options) BindFlags(fs *bootflag.FlagSet)  {
	if o.CustomBindFlagsFunc != nil {
		o.CustomBindFlagsFunc(fs)
		return
	}


	fs.StringVar(
		&o.Network,
		utils.BuildFlagName(defaultGoRedisFlagsPrefix, "network"),
		defaultGoRedisNetwork,
		"the network type, either tcp or unix")
}

func (o *Options) Parse() (err error) {
	if o.CustomParseFunc != nil {
		return o.CustomParseFunc()
	}
	return nil
}


