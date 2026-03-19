package service

import (
	"homework4/internal/model"
	"homework4/internal/repository"
	"homework4/pkg/apperror"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
	postRepo    *repository.PostRepository
}

func NewCommentService(commentRepo *repository.CommentRepository, postRepo *repository.PostRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func (s *CommentService) Create(postID, authorID uint, content string) (*model.Comment, error) {
	if _, err := s.postRepo.FindByID(postID); err != nil {
		return nil, apperror.NotFound("未找到文章", err)
	}

	comment := &model.Comment{
		Content:  content,
		PostID:   postID,
		AuthorID: authorID,
	}
	if err := s.commentRepo.Create(comment); err != nil {
		return nil, apperror.Internal("新增评论失败", err)
	}

	comments, err := s.commentRepo.ListByPostID(postID)
	if err != nil {
		return nil, apperror.Internal("加载评论失败", err)
	}
	for _, c := range comments {
		if c.ID == comment.ID {
			cp := c
			return &cp, nil
		}
	}
	return comment, nil
}

func (s *CommentService) ListByPostID(postID uint) ([]model.Comment, error) {
	if _, err := s.postRepo.FindByID(postID); err != nil {
		return nil, apperror.NotFound("文章未找到", err)
	}
	comments, err := s.commentRepo.ListByPostID(postID)
	if err != nil {
		return nil, apperror.Internal("加载评论失败", err)
	}
	return comments, nil
}
