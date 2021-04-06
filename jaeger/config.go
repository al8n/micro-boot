package jaeger

import (
	bootflag "github.com/al8n/micro-boot/flag"
	"github.com/al8n/micro-boot/internal/utils"
	"github.com/opentracing/opentracing-go"
)

var (
	defaultJaegerFlagsPrefix = "jaeger"
)

type Config struct {
	// ServiceName specifies the service name to use on the tracer.
	// Can be provided by FromEnv() via the environment variable named JAEGER_SERVICE_NAME
	ServiceName string `yaml:"service-name" json:"service-name"`

	// Disabled can be provided by FromEnv() via the environment variable named JAEGER_DISABLED
	Disabled bool `yaml:"disabled" json:"disabled"`

	// RPCMetrics can be provided by FromEnv() via the environment variable named JAEGER_RPC_METRICS
	RPCMetrics bool `yaml:"rpc-metrics" json:"rpc-metrics"`

	// Tags can be provided by FromEnv() via the environment variable named JAEGER_TAGS
	Tags []opentracing.Tag `yaml:"tags" json:"tags"`

	Sampler *SamplerConfig `yaml:"sampler" json:"sampler"`


	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-" mapstructure:",omitempty"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-" mapstructure:",omitempty"`
}


func (c *Config) BindFlags(fs *bootflag.FlagSet)  {


	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}


	fs.StringVar(
		&c.ServiceName,
		utils.BuildFlagName(defaultJaegerFlagsPrefix, "service-name"),
		"",
		"specifies the service name to use on the tracer")

	fs.BoolVar(
		&c.Disabled,
		utils.BuildFlagName(defaultJaegerFlagsPrefix, "disabled"),
		false,
		"whether the tracer is disabled or not")

	fs.BoolVar(
		&c.RPCMetrics,
		utils.BuildFlagName(defaultJaegerFlagsPrefix, "rpc-metrics"),
		false,
		"whether to store RPC metrics")

	c.Sampler.BindFlags(fs)
}

func (c *Config) Parse() (err error) {
	if c.Sampler != nil {
		if err = c.Sampler.Parse(); err != nil {
			return err
		}
	}

	if c.CustomParseFunc != nil {
		return c.CustomParseFunc()
	}
	return nil
}


