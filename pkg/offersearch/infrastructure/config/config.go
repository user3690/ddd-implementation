package config

type Config struct {
	DaBaseUrl string
}

func NewConfig() Config {
	return Config{
		DaBaseUrl: "https://api.direktanbindung.de",
	}
}
