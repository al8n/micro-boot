package flag

import (
	"errors"
	"fmt"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"strings"
)

type generator int

const (
	Random64 generator = iota
	Random128
	RandomTimestamped
)

const errGeneratorStr = "invalid string being converted to idgenerator.IDGenerator"

type extractGeneratorValue struct {
	value *idgenerator.IDGenerator
	typ *generator
}

func newZipkinGeneratorValue(val idgenerator.IDGenerator, p *idgenerator.IDGenerator) *extractGeneratorValue {
	ssv := new(extractGeneratorValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (e extractGeneratorValue) Type() string {
	return "zipkinIDGenerator"
}

func (e extractGeneratorValue) String() (result string) {
	var policy = map[generator]string{
		Random64: "ExtractFailurePolicyRestart",
		Random128: "ExtractFailurePolicyError",
		RandomTimestamped: "ExtractFailurePolicyTagAndRestart",
	}
	return fmt.Sprintf("error handling policy: %s", policy[*e.typ])
}

func (e *extractGeneratorValue) Set(val string) (err error)  {
	value := strings.ToLower(dashBlankReplacer.Replace(val))
	if strings.Contains(value, "64") {
		*e.value = idgenerator.NewRandom64()
		*e.typ = Random64
		return nil
	} else if strings.Contains(value, "128") {
		*e.value = idgenerator.NewRandom128()
		*e.typ = Random128
		return nil
	} else if strings.Contains(value, "timestamp") {
		*e.value = idgenerator.NewRandomTimestamped()
		*e.typ=RandomTimestamped
		return nil
	} else {
		return errors.New("unsupported zipkin id generator")
	}
}

func extractGeneratorConv(val string) (interface{}, error) {
	value := strings.ToLower(dashBlankReplacer.Replace(val))
	if strings.Contains(value, "64") {
		return idgenerator.NewRandom64(), nil
	} else if strings.Contains(value, "128") {
		return idgenerator.NewRandom128(), nil
	} else if strings.Contains(value, "timestamp") {
		return idgenerator.NewRandomTimestamped(), nil
	} else {
		return nil, fmt.Errorf("%s: %s", errGeneratorStr, val)
	}
}

// GetZipkinIDGenerator return the idgenerator.IDGenerator value of a flag with the given name
func (f *FlagSet) GetZipkinIDGenerator(name string) (idgenerator.IDGenerator, error) {
	val, err := f.getFlagType(name, "zipkinIDGenerator", extractGeneratorConv)
	if err != nil {
		return nil, err
	}
	return val.(idgenerator.IDGenerator), nil
}

// ZipkinIDGeneratorVar defines a idgenerator.IDGenerator flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) ZipkinIDGeneratorVar(p *idgenerator.IDGenerator, name string, value idgenerator.IDGenerator, usage string) {
	f.fs.VarP(newZipkinGeneratorValue(value, p), name, "", usage)
}

// ZipkinIDGeneratorVarP is like ZipkinIDGeneratorVarVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) ZipkinIDGeneratorVarP(p *idgenerator.IDGenerator, name, shorthand string, value idgenerator.IDGenerator, usage string) {
	f.fs.VarP(newZipkinGeneratorValue(value, p), name, shorthand, usage)
}


// ZipkinExtractFailurePolicy defines a zipkin.ExtractFailurePolicy flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) ZipkinIDGenerator(name string, value idgenerator.IDGenerator, usage string) *idgenerator.IDGenerator {
	p := new(idgenerator.IDGenerator)
	f.ZipkinIDGeneratorVarP(p, name, "", value, usage)
	return p
}


// ZipkinExtractFailurePolicyP is like ZipkinExtractFailurePolicy, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) ZipkinIDGeneratorP(name, shorthand string, value idgenerator.IDGenerator, usage string) *idgenerator.IDGenerator {
	p := new(idgenerator.IDGenerator)
	f.ZipkinIDGeneratorVarP(p, name, shorthand, value, usage)
	return p
}