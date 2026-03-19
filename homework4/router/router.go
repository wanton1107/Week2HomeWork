package router

import (
	"homework4/config"
	"homework4/internal/handler"
	"homework4/internal/middleware"
	"homework4/internal/repository"
	"homework4/internal/service"
	"homework4/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, log *logrus.Logger, db *gorm.DB) *gin.Engine {
	gin.SetMode(cfg.Server.Mode)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestLogger(log))
	r.Use(middleware.ErrorHandler(log))

	jwtMgr := jwt.NewManager(cfg.JWT.Secret, cfg.JWT.Issuer, cfg.JWT.ExpireHour)

	// 初始化持久层
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	// 初始化Service
	authService := service.NewAuthService(userRepo, jwtMgr)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRepo, postRepo)

	// 初始化Handler
	authHandler := handler.NewAuthHandler(authService)
	postHandler := handler.NewPostHandler(postService)
	commentHandler := handler.NewCommentHandler(commentService)

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		api.GET("/posts", postHandler.ListPost)
		api.GET("/posts/:id", postHandler.GetPost)
		api.GET("/posts/:id/comments", commentHandler.ListComment)

		authRequired := api.Group("")
		authRequired.Use(middleware.AuthMiddleware(jwtMgr))
		{
			authRequired.POST("/posts", postHandler.CreatePost)
			authRequired.PUT("/posts/:id", postHandler.UpdatePost)
			authRequired.DELETE("/posts/:id", postHandler.DeletePost)
			authRequired.POST("/posts/:id/comments", commentHandler.CreateComment)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
