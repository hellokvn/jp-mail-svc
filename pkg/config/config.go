package config

import "github.com/spf13/viper"

type Config struct {
	DBUrl        string `mapstructure:"DB_URL"`
	AMQPUrl      string `mapstructure:"AMQP_URL"`
	MailHost     string `mapstructure:"MAIL_HOST"`
	MailUser     string `mapstructure:"MAIL_USER"`
	MailPassword string `mapstructure:"MAIL_PASSWORD"`
	MailPort     string `mapstructure:"MAIL_PORT"`
	MailFrom     string `mapstructure:"MAIL_FROM"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
