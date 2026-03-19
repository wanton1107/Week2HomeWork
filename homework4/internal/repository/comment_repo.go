package repository

import (
	"homework4/internal/model"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) ListByPostID(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	if err := r.db.Preload("Author").Where("post_id = ?", postID).Order("id asc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
