package mysql

import (
	"fmt"

	"github.com/Sparshith/songski/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB(config *config.MySQL) (*gorm.DB, error) {
	configStr := getConfigString(config)
	db, err := gorm.Open("mysql", configStr)
	if config.Debug {
		db.LogMode(true)
	}
	return db, err
}

func getConfigString(config *config.MySQL) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
