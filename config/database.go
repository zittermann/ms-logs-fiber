package config

import (
	"fmt"
	"ms_logs_go/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DATABASE_FILE = "config.yml"

func CreateConnection() *gorm.DB {

	config := LoadConfig(DATABASE_FILE)

	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?parseTime=True",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	helper.ErrorPanic(err, `No se pudo establecer 
		la conexión con la base de datos...`)

	return db

}
