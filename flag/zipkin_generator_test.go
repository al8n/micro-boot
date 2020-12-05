package flag

import (
	"github.com/openzipkin/zipkin-go/idgenerator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setUpZipkinGeneratorFlagSet(g *idgenerator.IDGenerator) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.ZipkinIDGeneratorVar(g, "generator", Random64, "Command zipkin generator!")
	return f
}

func setUpZipkinGeneratorWithDefault(g *idgenerator.IDGenerator) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.ZipkinIDGeneratorVarP(g, "generator", "g", RandomTimestamped, "Command zipkin generator!")
	return f
}

func TestEmptyZipkinGenerator(t *testing.T) {
	var g idgenerator.IDGenerator
	f := setUpZipkinGeneratorFlagSet(&g)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getG, err := f.GetZipkinIDGenerator("generator")
	assert.NoError(t, err)

	assert.Equal(t, getG, idgenerator.NewRandom64())
}

func TestEmptyZipkinGeneratorWithDefault(t *testing.T) {
	var g idgenerator.IDGenerator
	f := setUpZipkinGeneratorWithDefault(&g)
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getG, err := f.GetZipkinIDGenerator("generator")
	assert.NoError(t, err)

	assert.Equal(t, getG, idgenerator.NewRandomTimestamped())
}

func TestZipkinGenerator64(t *testing.T) {
	var g *idgenerator.IDGenerator
	f := NewFlagSet("test", ContinueOnError)
	g = f.ZipkinIDGenerator("generator", Random128, "")
	err := f.fs.Parse([]string{"--generator=64"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	assert.Equal(t, idgenerator.NewRandom64(), *g)
}

func TestZipkinGenerator128(t *testing.T) {
	var g *idgenerator.IDGenerator
	f := NewFlagSet("test", ContinueOnError)
	g = f.ZipkinIDGenerator("generator", Random64, "")
	err := f.fs.Parse([]string{"--generator=128"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	assert.Equal(t, idgenerator.NewRandom128(), *g)
}

func TestZipkinGeneratorTimeStamped(t *testing.T) {
	var g *idgenerator.IDGenerator
	f := NewFlagSet("test", ContinueOnError)
	g = f.ZipkinIDGeneratorP("generator", "g", Random128, "")
	err := f.fs.Parse([]string{"-g timestamped"})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getG, err := f.GetZipkinIDGenerator("generator")
	assert.NoError(t, err)

	assert.Equal(t, getG, *g)
}

func TestZipkinGeneratorSetError(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.ZipkinIDGenerator("generator", Random64, "")
	err := f.fs.Parse([]string{"--generator=error"})
	assert.Error(t, err)
}

func TestZipkinGeneratorError(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.ZipkinIDGenerator("generator", -2, "")
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getG, err := f.GetZipkinIDGenerator("generator")
	assert.NoError(t, err)
	assert.Equal(t, getG, idgenerator.NewRandom64())
}