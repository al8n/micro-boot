package flag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIp4PortIntValue_Set(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortInt("ipv4", 8080, "")

	err := f.fs.Parse([]string{"--ipv4=9090"})
	assert.NoError(t, err)


	getIPv4, err := f.GetIPv4PortInt("ipv4")
	assert.NoError(t, err)
	assert.Equal(t, 9090, getIPv4)
}

func TestIp4PortIntValue(t *testing.T) {
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortIntP("ipv4", "i", 8080, "")

	err := f.fs.Parse([]string{"-i 9090"})
	assert.NoError(t, err)


	getIPv4, err := f.GetIPv4PortInt("ipv4")
	assert.NoError(t, err)
	assert.Equal(t,  9090, getIPv4)
}

func TestIp4PortIntValueVarP(t *testing.T) {
	var ipv4 int
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortIntVarP(&ipv4, "ipv4", "i", 8080, "")

	err := f.fs.Parse([]string{"-i 9090"})
	assert.NoError(t, err)
	assert.Equal(t,  9090, ipv4)
}

func TestIp4PortIntValueVar(t *testing.T) {
	var ipv4 int
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortIntVar(&ipv4, "ipv4",8080, "")

	err := f.fs.Parse([]string{"--ipv4=9090"})
	assert.NoError(t, err)
	assert.Equal(t,  9090, ipv4)
}

func TestIp4PortIntValueVarError(t *testing.T) {
	var ipv4 int
	f := NewFlagSet("test", ContinueOnError)
	f.IPv4PortIntVar(&ipv4, "ipv4",8080, "")

	err := f.fs.Parse([]string{"--ipv4=90900"})
	assert.Error(t, err)
}
