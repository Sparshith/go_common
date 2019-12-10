package config

import (
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

/*
Data stores
*/

type MySQL struct {
	Database string
	Username string
	Password string
	Host     string
	Port     string
	Debug    bool
}

/*
External Applications
*/
type Spotify struct {
	ClientId       string
	ClientSecretId string
}

type Config struct {
	Spotify *Spotify
	MySQL   *MySQL
}

type RedigoConfig struct {
	Address         string
	MaxIdle         int
	IdleTimeout     time.Duration
	Wait            bool
}

func Initialize() *Config {
	setUpViper("./configs/", GetEnv())
	config := &Config{}
	config.MySQL = initializeMySQL()
	config.Spotify = initializeSpotify()
	return config
}

func initializeMySQL() *MySQL {
	return &MySQL{
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		Database: viper.GetString("mysql.database"),
		Debug:    viper.GetBool("mysql.debug"),
	}
}

func initializeSpotify() *Spotify {
	return &Spotify{
		ClientId:       viper.GetString("spotify.client_id"),
		ClientSecretId: viper.GetString("spotify.client_secret_id"),
	}
}

func setUpViper(directory, filename string) {
	viper.AddConfigPath(directory)
	viper.SetConfigName(filename)
	viper.ReadInConfig()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
}

func GetEnv() string {
	return os.Getenv("GO_ENV")
}
