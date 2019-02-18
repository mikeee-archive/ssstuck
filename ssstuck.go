package ssstuck

import (
	"fmt"
)

// Config is the basic structure to hold required startup variables
type Config struct {
	Port int
}

// CheckConfig validates the input config
func CheckConfig(config Config) error {
	if config.Port < 1 || config.Port > 65535 {
		return fmt.Errorf("port(%d) please specify a port between 1-65535", config.Port)
	}
	return nil
}
