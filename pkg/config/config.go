package config

import "github.com/spf13/viper"

type Config struct {
	vp *viper.Viper
}

var config *Config

// Init creates a new Config instance.
func Init(configPath string) {
	New(configPath)
}

// New creates a new Config instance.
func New(configPath string) *Config {
	if config == nil {
		config = &Config{
			vp: viper.New(),
		}
		config.vp.SetConfigFile(configPath)
		err := config.vp.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	return config
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
