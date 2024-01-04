package config

type Options interface {
	apply(c *Config)
}

type optionFunc func(c *Config)

func (f optionFunc) apply(c *Config) {
	f(c)
}

// WithAutomaticEnv enables the automatic env variables binding.
func WithAutomaticEnv() Options {
	return optionFunc(func(c *Config) {
		c.vp.AutomaticEnv()
	})
}

// WithEnvPrefix sets the prefix for environment variables binding.
func WithEnvPrefix(prefix string) Options {
	return optionFunc(func(c *Config) {
		c.vp.SetEnvPrefix(prefix)
	})
}

// WithDefaultValues sets the default values for the config.
func WithDefaultValues(defaults map[string]interface{}) Options {
	return optionFunc(func(c *Config) {
		for key, value := range defaults {
			c.vp.SetDefault(key, value)
		}
	})
}
