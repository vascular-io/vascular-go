package vascular

import "flag"

type Credentials struct {
	ApiKey *string
	AppKey *string
}

type Config struct {
	Credentials *Credentials
	Port        *string
	UserID      *string
}

var addr = flag.String("addr", "api.vascular.io", "the address to connect to")

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithCredentials(creds *Credentials) *Config {
	c.Credentials = creds
	return c
}

func (c *Config) WithUserID(userID string) *Config {
	c.UserID = &userID
	return c
}
