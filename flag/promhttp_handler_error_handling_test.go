package flag

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerErrorHandlingHTTP(t *testing.T) {
	var p *promhttp.HandlerErrorHandling
	f := NewFlagSet("test", ContinueOnError)
	p = f.PrometheusHandlerErrorHandling("policy", promhttp.ContinueOnError, "")
	err := f.fs.Parse([]string{"--policy=http"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetPrometheusHandlerErrorHandling("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, *p)
}

func TestHandlerErrorHandlingContinue(t *testing.T) {
	var p *promhttp.HandlerErrorHandling
	f := NewFlagSet("test", ContinueOnError)
	p = f.PrometheusHandlerErrorHandlingP("policy", "p",promhttp.ContinueOnError, "")
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetPrometheusHandlerErrorHandling("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, *p)
}

func TestHandlerErrorHandlingPanic(t *testing.T) {
	var p promhttp.HandlerErrorHandling
	f := NewFlagSet("test", ContinueOnError)
	f.PrometheusHandlerErrorHandlingVarP(&p, "policy", "p", promhttp.ContinueOnError, "")
	err := f.fs.Parse([]string{"-p panic"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetPrometheusHandlerErrorHandling("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, p)
}

func TestHandlerErrorHandling(t *testing.T) {
	var p promhttp.HandlerErrorHandling
	f := NewFlagSet("test", ContinueOnError)
	f.PrometheusHandlerErrorHandlingVar(&p, "policy", promhttp.HTTPErrorOnError, "")
	err := f.fs.Parse([]string{"--policy=continue"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetPrometheusHandlerErrorHandling("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, p)
}

func TestHandlerErrorHandlingError(t *testing.T) {
	var p promhttp.HandlerErrorHandling
	f := NewFlagSet("test", ContinueOnError)
	f.PrometheusHandlerErrorHandlingVar(&p, "policy", -6, "")
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	_, err = f.GetPrometheusHandlerErrorHandling("policy")
	assert.Error(t, err)
}