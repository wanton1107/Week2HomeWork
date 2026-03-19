package service

import (
	"errors"
	"homework4/internal/model"
	"homework4/internal/repository"
	"homework4/pkg/apperror"

	"gorm.io/gorm"
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) Create(title, content string, authorID uint) (*model.Post, error) {
	post := &model.Post{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}
	if err := s.postRepo.Create(post); err != nil {
		return nil, apperror.Internal("创建文章失败", err)
	}
	return s.postRepo.FindByID(post.ID)
}

func (s *PostService) List() ([]model.Post, error) {
	posts, err := s.postRepo.FindAll()
	if err != nil {
		return nil, apperror.Internal("获取文章列表失败", err)
	}
	return posts, nil
}

func (s *PostService) Detail(id uint) (*model.Post, error) {
	post, err := s.postRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("文章不存在", err)
		}
		return nil, apperror.Internal("加载文章内容失败", err)
	}
	return post, nil
}

func (s *PostService) Update(id uint, title, content string, currentUserID uint) (*model.Post, error) {
	post, err := s.postRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("文章不存在", err)
		}
		return nil, apperror.Internal("未找到文章", err)
	}
	if post.AuthorID != currentUserID {
		return nil, apperror.Forbidden("无权限修改该文章", nil)
	}

	post.Title = title
	post.Content = content
	if err := s.postRepo.Update(post); err != nil {
		return nil, apperror.Internal("更新文章内容失败", err)
	}
	return s.postRepo.FindByID(post.ID)
}

func (s *PostService) Delete(id uint, currentUserID uint) error {
	post, err := s.postRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("文章不存在", err)
		}
		return apperror.Internal("未找到文章", err)
	}
	if post.AuthorID != currentUserID {
		return apperror.Forbidden("无权限修改该文章", nil)
	}
	if err := s.postRepo.Delete(id); err != nil {
		return apperror.Internal("删除文章失败", err)
	}
	return nil
}
