package main

import (
	"fmt"
	"homework4/config"
	_ "homework4/docs"
	"homework4/internal/model"
	"homework4/internal/repository"
	"homework4/pkg/logger"
	"homework4/router"
)

// @title 项目接口
// @version 1.0 版本
// @description 博客项目API接口文档
// @host http://localhost:8080
// @BasePath /api/v1
func main() {
	cfg, err := config.InitConfig(".")
	if err != nil {
		panic(err)
	}

	log := logger.NewLogger(cfg.Log.Level)

	if err := repository.InitDB(cfg, log); err != nil {
		log.WithError(err).Fatal("failed to initialize database")
		return
	}

	if err := repository.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		log.WithError(err).Fatal("failed to run database migration")
		return
	}

	r := router.SetupRouter(cfg, log, repository.DB)
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Infof("server started at %s", addr)
	if err := r.Run(addr); err != nil {
		log.WithError(err).Fatal("failed to start http server")
		return
	}
}
