package config

type (
	Config struct {
		DB `yaml:"datasource"`
	}

	DB struct {
		URL string `yaml:"url"`
	}

	Log struct {
		Level string
	}
)

func Load(path string) (*Config, error) {
	return nil, nil
}
