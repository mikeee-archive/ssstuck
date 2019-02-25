package ssstuck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckConfig(t *testing.T) {
	tests := []struct {
		name     string
		value    Config
		expected assert.ErrorAssertionFunc
	}{
		{"0 is an invalid port", Config{Port: 0}, assert.Error},
		{"1 is a valid port", Config{Port: 1}, assert.NoError},
		{"22 is a valid port", Config{Port: 22}, assert.NoError},
		{"2222 is a valid port", Config{Port: 2222}, assert.NoError},
		{"65536 is an invalid port", Config{Port: 65536}, assert.Error},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.expected(t, CheckConfig(test.value))
		})
	}
}
