package service

import (
	"errors"
	"homework4/internal/model"
	"homework4/internal/repository"
	"homework4/pkg/apperror"
	"homework4/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repository.UserRepository
	jwtMgr   *jwt.Manager
}

func NewAuthService(userRepo *repository.UserRepository, jwtMgr *jwt.Manager) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtMgr:   jwtMgr,
	}
}

func (s *AuthService) Register(username, password string) error {
	_, err := s.userRepo.FindByUsername(username)
	if err == nil {
		return apperror.BadRequest("已存在的用户账号", nil)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.Internal("密码加密失败", err)
	}

	user := &model.User{
		Username: username,
		Password: string(hash),
	}
	if err := s.userRepo.Create(user); err != nil {
		return apperror.Internal("创建用户失败", err)
	}
	return nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", apperror.Unauthorized("用户名或密码错误", err)
		}
		return "", apperror.Internal("该用户不存在", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", apperror.Unauthorized("用户名或密码错误", err)
	}

	token, err := s.jwtMgr.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", apperror.Internal("token生成失败", err)
	}
	return token, nil
}
