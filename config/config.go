package config

type Config struct {
	Listen struct {
		Addr string
		Bind string
	}

	DEVMODE bool `yaml:"dev_mode"`

	Redis RedisConfig `yaml:"redis"`

	LogPath string `yaml:"log_path"`

	Code struct {
		Len int
		Expire int
		Width int
		Height int
	}
}

type RedisConfig struct {
	Addr string
	DB int
	Password string
}