package config

import (
	"os"
	"reflect"
	"testing"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func Test_initializeSpotify(t *testing.T) {
	os.Setenv("SPOTIFY_CLIENT_ID", "sparshith")
	os.Setenv("SPOTIFY_CLIENT_SECRET_ID", "nixie")
	Initialize()
	spotifyConfig := &Spotify{
		ClientId:       "sparshith",
		ClientSecretId: "nixie",
	}
	tests := []struct {
		name string
		want *Spotify
	}{
		{
			name: "Get spotify config",
			want: spotifyConfig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeSpotify(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeSpotify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Should return test as the env",
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initializeMySQL(t *testing.T) {
	os.Setenv("MYSQL_DATABASE", "dummy")
	os.Setenv("MYSQL_USERNAME", "nixie")
	os.Setenv("MYSQL_PASSWORD", "sparshith")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_HOST", "localhost")

	Initialize()
	mysqlConfig := &MySQL{
		Database: "dummy",
		Username: "nixie",
		Password: "sparshith",
		Port: "3306",
		Host: "localhost",
	}

	tests := []struct {
		name string
		want *MySQL
	}{
		{
			name: "Should fetch mysql config",
			want: mysqlConfig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeMySQL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeMySQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
