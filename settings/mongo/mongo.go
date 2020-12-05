package mongo

import (
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"time"
)

var (
	defaultMongoFlagsPrefix = "mongo"
)

func SetDefaultMongoFlagsPrefix(val string)  {
	defaultMongoFlagsPrefix = val
}

type ClientOptions struct {
	AppName                  *string                `json:"app-name" yaml:"app-name"`
	Auth                     *Credential 			`json:"auth" yaml:"auth"`

	//AutoEncryptionOptions    *AutoEncryptionOptions  `json:"auto-encryption" yaml:"auto-encryption"`

	ConnectTimeout           *time.Duration    `json:"connect-timeout" yaml:"connect-timeout"`
	Compressors              []string			`json:"compressors" yaml:"compressors"`
	Collection 				string        `json:"collection" yaml:"collection"`
	DB         				string        `json:"database" yaml:"database"`
	Direct                   *bool				`json:"direct" yaml:"direct"`
	DisableOCSPEndpointCheck *bool				`json:"disable-ocsp-endpoint-check" yaml:"disable-ocsp-endpoint-check"`
	HeartbeatInterval        *time.Duration     `json:"heartbeat-interval" yaml:"heartbeat-interval"`
	Hosts 					[]string 			`json:"hosts" yaml:"hosts"`
	LocalThreshold           *time.Duration    `json:"local-threshold" yaml:"local-threshold"`
	MaxConnIdleTime          *time.Duration    `json:"max-conn-idle-time" yaml:"max-conn-idle-time"`
	MaxPoolSize              *uint64 			`json:"max-pool-size" yaml:"max-pool-size"`
	MinPoolSize              *uint64 			`json:"min-pool-size" yaml:"min-pool-size"`

	// ReadConcern for replica sets and replica set shards determines which data to return from a query.
	ReadConcern              *ReadConcern        `json:"read-concern" yaml:"read-concern"`

	// ReadPreference determines which servers are considered suitable for read operations.
	ReadPreference           *ReadPref      		`json:"read-preference" yaml:"read-preference"`

	ReplicaSet               *string  			`json:"replica-set" yaml:"replica-set"`

	RetryReads               *bool  			`json:"retry-reads" yaml:"retry-reads"`

	RetryWrites              *bool 				`json:"retry-writes" yaml:"retry-writes"`

	ServerSelectionTimeout   *time.Duration     `json:"server-selection-timeout" yaml:"server-selection-timeout"`

	SocketTimeout            *time.Duration      `json:"socket-timeout" yaml:"socket-timeout"`

	URI 					*string 				`json:"uri" yaml:"uri"`

	ZlibLevel                *int 				`json:"zlib-level" yaml:"zlib-yaml"`

	ZstdLevel                *int 				`json:"zstd-level" yaml:"zstd-yaml"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (m *ClientOptions) BindFlags(fs *bootflag.FlagSet) {
	if m.CustomBindFlagsFunc != nil {
		m.CustomBindFlagsFunc(fs)
		return
	}

	m.Auth.BindFlags(fs)
	m.ReadConcern.BindFlags(fs)
	m.ReadPreference.BindFlags(fs)
}

func (m *ClientOptions) Parse() (err error) {
	if m.URI != nil {
		_, err = connstring.Parse(*m.URI)
		if err != nil {
			return fmt.Errorf("%s is not a valid MongoDB URI", *m.URI)
		}
	}

	return nil
}