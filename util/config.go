package util

import (
	"time"

	"github.com/spf13/viper"
)

// In order to get the value of the variables and store them in this struct, we need to use the unmarshaling feature of Viper.
// Viper uses the mapstructure package under the hood for unmarshaling values, so we use the mapstructure tags to specify the name of each config field.
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBUrl               string        `mapstructure:"DB_URL"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	SecretKey           string        `mapstructure:"SECRET_KEY"`
	SecretKeyAdmin      string        `mapstructure:"SECRET_KEY_ADMIN"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	MaxFileSize         int64         `mapstructure:"MAX_FILE_SIZE"`
	AwsRegion           string        `mapstructure:"AWS_REGION"`
	AwsAccessKeyId      string        `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey  string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	EmailSenderName     string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
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
