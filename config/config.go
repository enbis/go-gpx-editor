package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetDefault("Closure", "</trkseg>\n</trk>\n</gpx>")
	viper.SetDefault("Prefix", "<trkpt ")

	viper.AddConfigPath("./")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		log.Debugf("Using config file: %s", viper.ConfigFileUsed())
	}

	return nil
}
