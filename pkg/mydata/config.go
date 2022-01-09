package mydata

import (
	"time"

	"github.com/ppapapetrou76/go-utils/pkg/multierror"
	"github.com/ppapapetrou76/go-utils/pkg/validation"
)

// Config holds the required and optional configuration for the mydata client.
type Config struct {
	Host            string
	UserID          string
	SubscriptionKey string
	Timeout         time.Duration
}

// validate checks that the config's required attribute have values.
func (c *Config) validate() error {
	return multierror.NewPrefixed("api config validation",
		validation.IsRequired(c.SubscriptionKey, "subscription key"),
		validation.IsRequired(c.UserID, "user id"),
	).ErrorOrNil()
}

// defaults assigns default values to the optional config attributes.
func (c *Config) defaults() {
	if c.Timeout.Nanoseconds() == 0 {
		c.Timeout = defaultTimeout
	}

	if c.Host == "" {
		c.Host = defaultHost
	}
}
