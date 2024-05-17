package env

import (
	"strings"

	"github.com/spf13/viper"
)

// Environment environment
type Environment struct {
	Config *Config `mapstructure:"CONFIG"`
}

// Config config
type Config struct {
	Release    bool   `mapstructure:"RELEASE"`
	ServerPort string `mapstructure:"SERVERPORT"`
	Database   struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		Name     string `mapstructure:"NAME"`
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
	} `mapstructure:"DATABASE"`
	Mongo struct {
		Host     string `mapstructure:"HOST"`
		Name     string `mapstructure:"NAME"`
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
	} `mapstructure:"MONGO"`
}

// Read init env
func Read(path string) (*Environment, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetConfigType("yml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	env := &Environment{}
	err := v.Unmarshal(&env)
	if err != nil {
		return nil, err
	}
	return env, nil
}
