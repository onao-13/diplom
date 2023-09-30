package config

type Config struct {
	Port string
}

func Dev() Config {
	return Config{"8080"}
}
