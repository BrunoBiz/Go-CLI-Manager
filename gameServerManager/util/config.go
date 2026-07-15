package util

import "github.com/spf13/viper"

// Stores all configs - Reads with VIPER
type Config struct {
	TMUXSessionName   string `mapstructure:"TMUX_SESSION_NAME"`
	GameStartFilePath string `mapstructure:"GAME_START_PATH"`
	GameServerDir     string `mapstructure:"GAME_DIR"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("mgr")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
