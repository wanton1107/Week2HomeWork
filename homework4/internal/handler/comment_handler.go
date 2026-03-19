package handler

import (
	"homework4/internal/dto"
	"homework4/internal/middleware"
	"homework4/internal/service"
	"homework4/pkg/apperror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

// CreateComment godoc
// @Summary 新增评论
// @Tags Comment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "post id"
// @Param request body dto.CreateCommentRequest true "create comment request"
// @Success 201 {object} model.Comment
// @Failure 404 {object} map[string]string
// @Router /posts/{id}/comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID, err := parseID(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.BadRequest("请求参数异常", err))
		return
	}

	userIDValue, _ := c.Get(middleware.ContextUserIDKey)
	userID, ok := userIDValue.(uint)
	if !ok {
		c.Error(apperror.Unauthorized("未登录用户", nil))
		return
	}

	comment, err := h.commentService.Create(postID, userID, req.Content)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// ListComment godoc
// @Summary 获取指定文章的评论
// @Tags Comment
// @Produce json
// @Param id path int true "post id"
// @Success 200 {array} model.Comment
// @Failure 404 {object} map[string]string
// @Router /posts/{id}/comments [get]
func (h *CommentHandler) ListComment(c *gin.Context) {
	postID, err := parseID(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	comments, err := h.commentService.ListByPostID(postID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, comments)
}
