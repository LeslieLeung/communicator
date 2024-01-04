package database

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(dsn string) {
	db = New(dsn)
}

func New(dsn string) *gorm.DB {
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		panic(err)
	}
	// Ping to verify the connection
	conn, err := database.DB()
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	return database
}

func GetDB(c *gin.Context) *gorm.DB {
	return db.WithContext(c.Request.Context())
}
