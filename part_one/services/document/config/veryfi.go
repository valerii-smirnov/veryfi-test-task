package config

type VeryfiConfig struct {
	ClientID     string `env:"CLIENT_ID,required"`
	ClientSecret string `env:"CLIENT_SECRET,required"`
	Username     string `env:"USERNAME,required"`
	APIKey       string `env:"API_KEY,required"`
}
