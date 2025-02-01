package configs

// Config represents the main configuration structure for the application.
// It holds all the configuration settings organized in sub-structures.
// The configuration can be loaded from various sources using mapstructure.
type (
	Config struct {
		Service Service `mapstructure:"service"`
	}

	Service struct {
		Port string `mapstructure:"port"`
	}
)
