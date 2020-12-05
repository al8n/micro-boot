package mongo

import (
	"fmt"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/internal/utils"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/tag"
	"time"
)

const (
	defaultMaxStalenessUsage = "specify a maximum replication lag for reads from secondaries in a replica set"
	defaultHedgeEnabledUsage = "specifies whether or not hedged reads should be enabled in the server"
	defaultReadPrefModeUsage = "specify the read preference mode"
	defaultTagSetsUsage = "specify one or more read preference tags"
)

var (
	defaultReadPrefFlagsPrefix = defaultMongoFlagsPrefix + "-read-pref"
	defaultMaxStaleness = -1 * time.Millisecond
	defaultHedgeEnabled = false
	defaultTagSets = []tag.Set{}
	defaultReadPrefMode readpref.Mode = 0
)

func SetDefaultReadPrefFlagsPrefix(val string)  {
	defaultReadPrefFlagsPrefix = val
}

func SetDefaultReadPrefMaxStaleness(val time.Duration)  {
	defaultMaxStaleness = val
}

func SetDefaultReadPrefHedgeEnabled(val bool)  {
	defaultHedgeEnabled = val
}

func SetDefaultReadPrefMode(val readpref.Mode)  {
	defaultReadPrefMode = val
}

func SetDefaultReadPrefTagSets (val []tag.Set)  {
	defaultTagSets = val
}

// ReadPref determines which servers are considered suitable for read operations.
type ReadPref struct {
	MaxStaleness    time.Duration   `json:"max-staleness" yaml:"max-staleness"`
	Mode            readpref.Mode   `json:"mode" yaml:"mode"`
	TagSets         []tag.Set       `json:"tag-sets" yaml:"tag-sets"`
	HedgeEnabled    bool			`json:"hedge-enabled" yaml:"hedge-enabled"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure, if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure, if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`

	standardized bool
	srp         *readpref.ReadPref
}

func (rp *ReadPref) BindFlags(fs *bootflag.FlagSet) {
	if rp.CustomBindFlagsFunc != nil {
		rp.CustomBindFlagsFunc(fs)
		return
	}

	fs.DurationVar(
		&rp.MaxStaleness,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "max-staleness"), defaultMaxStaleness, defaultMaxStalenessUsage)

	fs.MongoReadPreferenceModeVar(&rp.Mode, utils.BuildFlagName(defaultReadConcernFlagsPrefix, "mode"), defaultReadPrefMode, defaultReadPrefModeUsage)

	fs.MongoTagSetSliceVar(&rp.TagSets, utils.BuildFlagName(defaultReadPrefFlagsPrefix, "tag-sets"), defaultTagSets, defaultTagSetsUsage)

	fs.BoolVar(
		&rp.HedgeEnabled,
		utils.BuildFlagName(defaultMongoFlagsPrefix, "hedge-enabled"), defaultHedgeEnabled, defaultHedgeEnabledUsage)
}

func (rp *ReadPref) Parse() (err error) {
	if rp.CustomParseFunc != nil {
		return rp.CustomParseFunc()
	}
	return nil
}

func (rp *ReadPref) Get() (srp *readpref.ReadPref, err error) {
	if !rp.standardized {
		return rp.Standardize()
	}

	return rp.srp, nil
}

func (rp *ReadPref) Standardize() (srp *readpref.ReadPref, err error) {
	if rp.Mode < 1 || rp.Mode > 5 {
		return nil, fmt.Errorf("no valid read preference mode is provided")
	}

	var opts []readpref.Option

	if rp.MaxStaleness != -1 * time.Millisecond {
		opts = append(opts, readpref.WithMaxStaleness(rp.MaxStaleness))
	}

	if len(rp.TagSets) > 0 {
		opts = append(opts, readpref.WithTagSets(rp.TagSets...))
	}

	if rp.HedgeEnabled {
		opts = append(opts, readpref.WithHedgeEnabled(rp.HedgeEnabled))
	}
	srp, err = readpref.New(rp.Mode, opts...)
	if err != nil {
		return nil, err
	}

	rp.standardized = true
	rp.srp = srp
	return
}
