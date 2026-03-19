package repository

import (
	"homework4/config"
	internalrepo "homework4/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config, log *logrus.Logger) error {
	if err := internalrepo.InitDB(cfg, log); err != nil {
		return err
	}
	DB = internalrepo.DB
	return nil
}
