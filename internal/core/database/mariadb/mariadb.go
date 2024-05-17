package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config config
type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Debug        bool
}

// Session session
type Session struct {
	Db *gorm.DB
}

// New new database connection
func NewMariaDB(c *Config) (*Session, error) {
	dns := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DatabaseName,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return &Session{
		Db: db,
	}, nil
}
