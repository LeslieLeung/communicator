package config

import "github.com/spf13/viper"

type Config struct {
	vp *viper.Viper
}

var config *Config

// Init creates a new Config instance.
func Init(options ...Options) {
	config = New(options...)
}

// New creates a new Config instance.
func New(options ...Options) *Config {
	cfg := &Config{
		vp: viper.New(),
	}
	for _, opt := range options {
		opt.apply(cfg)
	}
	return cfg
}

// GetConfig returns the Config instance.
func GetConfig() *Config {
	return config
}

func (c *Config) GetString(key string) string {
	return c.vp.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.vp.GetInt(key)
}

func (c *Config) GetBool(key string) bool {
	return c.vp.GetBool(key)
}
