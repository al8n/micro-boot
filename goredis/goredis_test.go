package goredis

import (
	"testing"
	"github.com/mitchellh/mapstructure"
)

func TestGoRedis(t *testing.T) {
	input := map[string]interface{}{
		"network": "tcp",
		"username":  "Al",
		"password": "12345678",
	}

	var opts Options
	err := mapstructure.Decode(input, &opts)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", opts)
}
