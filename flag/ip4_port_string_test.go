package flag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIp4PortStringValue_Set(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortString("ipv4", "8080", "")

	err := f.fs.Parse([]string{"--ipv4=9090"})
	assert.NoError(t, err)


	getIPv4, err := f.GetIPv4PortString("ipv4")
	assert.NoError(t, err)
	assert.Equal(t, "9090", getIPv4)
}

func TestIp4PortStringValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortStringP("ipv4", "i", "8080", "")

	err := f.fs.Parse([]string{"-i 9090"})
	assert.NoError(t, err)


	getIPv4, err := f.GetIPv4PortString("ipv4")
	assert.NoError(t, err)
	assert.Equal(t, "9090", getIPv4)
}

func TestIp4PortStringValueVarP(t *testing.T) {
	var ipv4 string
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortStringVarP(&ipv4, "ipv4", "i", "8080", "")

	err := f.fs.Parse([]string{"-i 9090"})
	assert.NoError(t, err)
	assert.Equal(t, "9090", ipv4)
}

func TestIp4PortStringValueVar(t *testing.T) {
	var ipv4 string
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortStringVar(&ipv4, "ipv4","8080", "")

	err := f.fs.Parse([]string{"--ipv4=9090"})
	assert.NoError(t, err)
	assert.Equal(t, "9090", ipv4)
}

func TestIp4PortStringValueVarError(t *testing.T) {
	var ipv4 string
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortStringVar(&ipv4, "ipv4","808000", "")

	err := f.fs.Parse([]string{"--ipv4=90900"})
	assert.Error(t, err)
}