package configs

type Config struct {
	BindAddr    string `toml:"bindAddr"`
	LogLevel    string `toml:"logLevel"`
	DatabaseURL string `toml:"databaseUrl`
	SecretKey   string `toml:"secretKey"`
	GrpcConAddr string `toml:"grpcConnectionAddr"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
