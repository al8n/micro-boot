package flag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIp4PortUintValue_Set(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortUint("ipv4", 8080, "")

	err := f.fs.Parse([]string{"--ipv4=9090"})
	assert.NoError(t, err)


	getIPv4, err := f.GetIPv4PortUint("ipv4")
	assert.NoError(t, err)
	assert.Equal(t, uint(9090), getIPv4)
}

func TestIp4PortUintValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortUintP("ipv4", "i", 8080, "")

	err := f.fs.Parse([]string{"-i 9090"})
	assert.NoError(t, err)


	getIPv4, err := f.GetIPv4PortUint("ipv4")
	assert.NoError(t, err)
	assert.Equal(t,  uint(9090), getIPv4)
}

func TestIp4PortUintValueVarP(t *testing.T) {
	var ipv4 uint
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortUintVarP(&ipv4, "ipv4", "i", 8080, "")

	err := f.fs.Parse([]string{"-i 9090"})
	assert.NoError(t, err)
	assert.Equal(t,  uint(9090), ipv4)
}

func TestIp4PortUintValueVar(t *testing.T) {
	var ipv4 uint
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortUintVar(&ipv4, "ipv4",8080, "")

	err := f.fs.Parse([]string{"--ipv4=9090"})
	assert.NoError(t, err)
	assert.Equal(t,  uint(9090), ipv4)
}

func TestIp4PortUintValueVarError(t *testing.T) {
	var ipv4 uint
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortUintVar(&ipv4, "ipv4",8080, "")

	err := f.fs.Parse([]string{"--ipv4=90900"})
	assert.Error(t, err)
}
