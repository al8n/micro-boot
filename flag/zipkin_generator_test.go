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

func TestEmptyZipkinGenerator128(t *testing.T) {
	var g *idgenerator.IDGenerator
	f := NewFlagSet("test", ContinueOnError)
	g = f.ZipkinIDGenerator("generator", Random128, "")
	err := f.fs.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getG, err := f.GetZipkinIDGenerator("generator")
	assert.NoError(t, err)

	assert.Equal(t, getG, *g)
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