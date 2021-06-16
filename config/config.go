package config

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigurationsWrapper struct {
	Redis       RedisConfigurations
	BasketRules BasketRules
}

type RedisConfigurations struct {
	ConnectionString string
	InstanceName     string
}

type BasketRules struct {
	MaximumItemsInBasket         int
	MaximumSameItemInBasket      int
	MaximumCharLimitOfBasketNote int
	ExpirationDurationInDays     int
}

type Configuration struct {
	ConfigurationsWrapper
}

func NewConfiguration() Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration ConfigurationsWrapper

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return Configuration{configuration}
}
