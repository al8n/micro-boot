package flag

import (
	"github.com/openzipkin/zipkin-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZipkinExtractFailurePolicyRestart(t *testing.T) {
	var p *zipkin.ExtractFailurePolicy
	f := NewFlagSet("test", ContinueOnError)
	p = f.ZipkinExtractFailurePolicy("policy", zipkin.ExtractFailurePolicyRestart, "")
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetZipkinExtractFailurePolicy("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, *p)
}

func TestZipkinExtractFailurePolicyOnError(t *testing.T) {
	var p *zipkin.ExtractFailurePolicy
	f := NewFlagSet("test", ContinueOnError)
	p = f.ZipkinExtractFailurePolicyP("policy", "p", zipkin.ExtractFailurePolicyError, "")
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetZipkinExtractFailurePolicy("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, *p)
}

func TestZipkinExtractFailurePolicyOnTag(t *testing.T) {
	var p zipkin.ExtractFailurePolicy
	f := NewFlagSet("test", ContinueOnError)
	f.ZipkinExtractFailurePolicyVarP(&p, "policy", "p", zipkin.ExtractFailurePolicyError, "")
	err := f.fs.Parse([]string{"-p tag"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetZipkinExtractFailurePolicy("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, p)
}

func TestZipkinExtractFailurePolicy(t *testing.T) {
	var p zipkin.ExtractFailurePolicy
	f := NewFlagSet("test", ContinueOnError)
	f.ZipkinExtractFailurePolicyVar(&p, "policy", zipkin.ExtractFailurePolicyError, "")
	err := f.fs.Parse([]string{"--policy=tag"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getP, err := f.GetZipkinExtractFailurePolicy("policy")
	assert.NoError(t, err)
	assert.Equal(t, getP, p)
}