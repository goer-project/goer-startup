package server

import (
	"path/filepath"
	"strings"

	"github.com/marmotedu/component-base/pkg/util/homedir"
	"github.com/marmotedu/log"
	"github.com/spf13/viper"
)

const (
	// RecommendedName defines the default project name
	RecommendedName = "goer"
)

var (
	// RecommendedHomeDir defines the default directory used to place all service configurations.
	RecommendedHomeDir = "." + RecommendedName

	// RecommendedEnvPrefix defines the ENV prefix used by all service.
	RecommendedEnvPrefix = strings.ToUpper(RecommendedName)
)

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Join(homedir.HomeDir(), RecommendedHomeDir))
		viper.AddConfigPath("/etc/" + RecommendedName)
		viper.SetConfigName(defaultName)
	}

	// Use config file from the flag.
	viper.SetConfigType("yaml")              // set the type of the configuration to yaml.
	viper.AutomaticEnv()                     // read in environment variables that match.
	viper.SetEnvPrefix(RecommendedEnvPrefix) // set ENVIRONMENT variables prefix to IAM.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("WARNING: viper failed to discover and load the configuration file: %s", err.Error())
	}
}
