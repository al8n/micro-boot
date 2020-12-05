package flag

import (
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strings"
)

const errHandlerErrorHandlingStr = "invalid string being converted to promhttp.HandlerErrorHandling"

type handlerErrorHandlingValue struct {
	value *promhttp.HandlerErrorHandling
} 

func newHandlerErrorHandlingValue(val promhttp.HandlerErrorHandling, p *promhttp.HandlerErrorHandling) *handlerErrorHandlingValue {
	ssv := new(handlerErrorHandlingValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (heh handlerErrorHandlingValue) Type() string {
	return "prometheusHandlerErrorHandling"
}

func (heh handlerErrorHandlingValue) String() (result string) {
	var policy = map[promhttp.HandlerErrorHandling]string{
		promhttp.HTTPErrorOnError: "HTTPErrorOnError",
		promhttp.ContinueOnError: "ContinueOnError",
		promhttp.PanicOnError: "PanicOnError",
	}
	return fmt.Sprintf("error handling policy: %s", policy[*heh.value])
}

func (heh *handlerErrorHandlingValue) Set(val string) (err error)  {
	value := strings.ToLower(dashBlankReplacer.Replace(val))
	if strings.Contains(value, "http") || strings.Contains(value, "erroronerror") {
		*heh.value = 0
		return nil
	} else if strings.Contains(value, "continue") {
		*heh.value = 1
		return nil
	} else if strings.Contains(value, "panic") {
		*heh.value = 2
		return nil
	} else {
		return errors.New("unsupported error handling policy")
	}
}

func handlerErrorHandlingConv(val string) (interface{}, error) {
	value := strings.ToLower(dashBlankReplacer.Replace(val))
	if strings.Contains(value, "http") || strings.Contains(value, "erroronerror") {
		return promhttp.HTTPErrorOnError, nil
	} else if strings.Contains(value, "continue") {
		return promhttp.ContinueOnError, nil
	} else if strings.Contains(value, "panic") {
		return promhttp.PanicOnError, nil
	} else {
		return nil, fmt.Errorf("%s: %s", errHandlerErrorHandlingStr, val)
	}
} 

// GetPrometheusHandlerErrorHandling return the promhttp.HandlerErrorHandling value of a flag with the given name
func (f *FlagSet) GetPrometheusHandlerErrorHandling(name string) (promhttp.HandlerErrorHandling, error) {
	val, err := f.getFlagType(name, "prometheusHandlerErrorHandling", handlerErrorHandlingConv)
	if err != nil {
		return 0, err
	}
	return val.(promhttp.HandlerErrorHandling), nil
}

// PrometheusHandlerErrorHandlingVar defines a promhttp.HandlerErrorHandling flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) PrometheusHandlerErrorHandlingVar(p *promhttp.HandlerErrorHandling, name string, value promhttp.HandlerErrorHandling, usage string) {
	f.fs.VarP(newHandlerErrorHandlingValue(value, p), name, "", usage)
}

// PrometheusHandlerErrorHandlingVarP is like PrometheusHandlerErrorHandlingVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) PrometheusHandlerErrorHandlingVarP(p *promhttp.HandlerErrorHandling, name, shorthand string, value promhttp.HandlerErrorHandling, usage string) {
	f.fs.VarP(newHandlerErrorHandlingValue(value, p), name, shorthand, usage)
}


// PrometheusHandlerErrorHandling defines a promhttp.HandlerErrorHandling flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) PrometheusHandlerErrorHandling(name string, value promhttp.HandlerErrorHandling, usage string) *promhttp.HandlerErrorHandling {
	p := new(promhttp.HandlerErrorHandling)
	f.PrometheusHandlerErrorHandlingVarP(p, name, "", value, usage)
	return p
}


// PrometheusHandlerErrorHandlingP is like PrometheusHandlerErrorHandling, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) PrometheusHandlerErrorHandlingP(name, shorthand string, value promhttp.HandlerErrorHandling, usage string) *promhttp.HandlerErrorHandling {
	p := new(promhttp.HandlerErrorHandling)
	f.PrometheusHandlerErrorHandlingVarP(p, name, shorthand, value, usage)
	return p
}