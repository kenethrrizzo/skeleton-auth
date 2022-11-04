package config

type Config struct {
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration string `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) {}
