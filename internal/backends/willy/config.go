package willy

import "time"

// Config defines the Willy backend configuration.
type Config struct {
	SubscriptionKey string        `mapstructure:"subscription_key"`
	RequestTimeout  time.Duration `mapstructure:"request_timeout"`
}
