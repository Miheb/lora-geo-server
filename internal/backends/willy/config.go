package willy

import "time"

// Config defines the Willy backend configuration.
type Config struct {
	RequestTimeout  time.Duration `mapstructure:"request_timeout"`
}
