package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:20;unique" json:"username"`
	Password  string    `gorm:"size:60" json:"password"`
	Email     string    `gorm:"size:30;not null" json:"email"`
	Posts     []Post    `json:"posts"`
	PostCount uint      `gorm:"default:0;not null" json:"post_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Post 文章
type Post struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"not null" json:"user_id"`
	Title         string    `gorm:"size:30" json:"title"`
	Content       string    `gorm:"type:text" json:"content"`
	Status        uint8     `gorm:"default:0;not null" json:"status"`
	User          User      `gorm:"foreignKey:UserID" json:"-"`
	CommentStatus uint8     `gorm:"default:0;not null" json:"comment_status"`
	Comments      []Comment `json:"comments"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (post *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", post.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (post *Post) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", post.UserID).UpdateColumn("post_count", gorm.Expr("post_count - ?", 1)).Error
}

// Comment 评论
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Content   string    `gorm:"type:text" json:"content"`
	Post      Post      `gorm:"foreignKey:PostID" json:"-"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (comment *Comment) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&Post{}).Where(&Post{ID: comment.PostID}).Updates(map[string]interface{}{"comment_status": 1}).Error
}

func (comment *Comment) AfterDelete(tx *gorm.DB) error {
	// 评论数
	var sum int64
	if err := tx.Model(&Comment{}).Where("post_id=?", comment.PostID).Count(&sum).Error; err != nil {
		return err
	}

	status := 0
	if sum != 0 {
		status = 1
	}
	if err := tx.Model(&Post{}).Where("id=?", comment.PostID).Updates(map[string]interface{}{"comment_status": status}).Error; err != nil {
		return err
	}
	return nil
}
