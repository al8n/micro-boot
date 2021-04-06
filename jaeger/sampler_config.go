package jaeger

import (
	bootflag "github.com/al8n/micro-boot/flag"
	"github.com/al8n/micro-boot/internal/utils"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	"time"
)

var (
	defaultJaegerSamplerFlagsPrefix = defaultJaegerFlagsPrefix + "-sampler"
)

type SamplerConfig struct {
	// Type specifies the type of the sampler: const, probabilistic, rateLimiting, or remote.
	// Can be provided by FromEnv() via the environment variable named JAEGER_SAMPLER_TYPE
	Type string `yaml:"type" json:"type"`

	// Param is a value passed to the sampler.
	// Valid values for Param field are:
	// - for "const" sampler, 0 or 1 for always false/true respectively
	// - for "probabilistic" sampler, a probability between 0 and 1
	// - for "rateLimiting" sampler, the number of spans per second
	// - for "remote" sampler, param is the same as for "probabilistic"
	//   and indicates the initial sampling rate before the actual one
	//   is received from the mothership.
	// Can be provided by FromEnv() via the environment variable named JAEGER_SAMPLER_PARAM
	Param float64 `yaml:"param" json:"param"`

	// SamplingServerURL is the URL of sampling manager that can provide
	// sampling strategy to this service.
	// Can be provided by FromEnv() via the environment variable named JAEGER_SAMPLING_ENDPOINT
	SamplingServerURL string `yaml:"sampling-server-url" json:"sampling-server-url"`

	// SamplingRefreshInterval controls how often the remotely controlled sampler will poll
	// sampling manager for the appropriate sampling strategy.
	// Can be provided by FromEnv() via the environment variable named JAEGER_SAMPLER_REFRESH_INTERVAL
	SamplingRefreshInterval time.Duration `yaml:"sampling-refresh-interval" json:"sampling-refresh-interval"`

	// MaxOperations is the maximum number of operations that the PerOperationSampler
	// will keep track of. If an operation is not tracked, a default probabilistic
	// sampler will be used rather than the per operation specific sampler.
	// Can be provided by FromEnv() via the environment variable named JAEGER_SAMPLER_MAX_OPERATIONS.
	MaxOperations int `yaml:"max-operations" json:"max-operations"`

	// Opt-in feature for applications that require late binding of span name via explicit
	// call to SetOperationName when using PerOperationSampler. When this feature is enabled,
	// the sampler will return retryable=true from OnCreateSpan(), thus leaving the sampling
	// decision as non-final (and the span as writeable). This may lead to degraded performance
	// in applications that always provide the correct span name on trace creation.
	//
	// For backwards compatibility this option is off by default.
	OperationNameLateBinding bool `yaml:"operation-name-late-binding" json:"operation-name-late-binding"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-" mapstructure:",omitempty"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-" mapstructure:",omitempty"`
}

func (sc *SamplerConfig) BindFlags(fs *bootflag.FlagSet)  {
	if sc.CustomBindFlagsFunc != nil {
		sc.CustomBindFlagsFunc(fs)
		return
	}


	fs.StringVar(
		&sc.Type,
		utils.BuildFlagName(defaultJaegerSamplerFlagsPrefix, "type"),
		"",
		"specifies the type of the sampler: const, probabilistic, rateLimiting, or remote")

	fs.Float64Var(
		&sc.Param,
		utils.BuildFlagName(defaultJaegerSamplerFlagsPrefix, "param"),
		0,
		"a value passed to the sampler")

	fs.StringVar(
		&sc.SamplingServerURL,
		utils.BuildFlagName(defaultJaegerSamplerFlagsPrefix, "sampling-server-url"),
		"",
		"the URL of sampling manager that can provide sampling strategy to this service")

	fs.DurationVar(
		&sc.SamplingRefreshInterval,
		utils.BuildFlagName(defaultJaegerSamplerFlagsPrefix, "sampling-refresh-interval"),
		time.Minute,
		"controls how often the remotely controlled sampler will poll sampling manager for the appropriate sampling strategy")

	fs.IntVar(
		&sc.MaxOperations,
		utils.BuildFlagName(defaultJaegerSamplerFlagsPrefix, "max-operations"),
		2000,
		"the maximum number of operations that the PerOperationSampler will keep track of")

	fs.BoolVar(
		&sc.OperationNameLateBinding,
		utils.BuildFlagName(defaultJaegerSamplerFlagsPrefix, "operation-name-late-binding"),
		false,
		"Opt-in feature for applications that require late binding of span name via explicit call to SetOperationName when using PerOperationSampler")
}

func (sc *SamplerConfig) Parse() (err error) {
	if sc.CustomParseFunc != nil {
		return sc.CustomParseFunc()
	}
	return nil
}

func (sc *SamplerConfig) Standardize() (config *jaegerconfig.SamplerConfig) {
	config = &jaegerconfig.SamplerConfig{
		Type:                     sc.Type,
		Param:                    sc.Param,
		SamplingServerURL:        sc.SamplingServerURL,
		SamplingRefreshInterval:  sc.SamplingRefreshInterval,
		MaxOperations:            sc.MaxOperations,
		OperationNameLateBinding: sc.OperationNameLateBinding,
	}
	return
}