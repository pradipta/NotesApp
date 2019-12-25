package dbconfig

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB = newConnection("root", "", "127.0.0.1")
