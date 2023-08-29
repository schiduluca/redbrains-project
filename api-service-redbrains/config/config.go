package config

type Config struct {
	Port              string `env:"PORT"`
	CryptoServicePATH string `env:"CRYPTO_SERVICE_PATH"`
}
