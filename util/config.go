package util

import "github.com/spf13/viper"

// In order to get the value of the variables and store them in this struct, we need to use the unmarshaling feature of Viper.
// Viper uses the mapstructure package under the hood for unmarshaling values, so we use the mapstructure tags to specify the name of each config field.
type Config struct {
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBUrl               string `mapstructure:"DB_URL"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	SecretKey           string `mapstructure:"SECRET_KEY"`
	AccessTokenDuration int    `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
