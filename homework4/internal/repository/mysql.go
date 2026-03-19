package repository

import (
	"fmt"
	"gorm.io/gorm/schema"
	"homework4/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(cfg *config.Config, log *logrus.Logger) error {
	db, err := gorm.Open(mysql.Open(cfg.Database.GetUrl()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "tb_",
		},
	})
	if err != nil {
		return fmt.Errorf("connect db failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get sql db failed: %w", err)
	}
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifeDuration())

	DB = db
	log.Info("database initialized")
	return nil
}
