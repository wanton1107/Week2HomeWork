package handler

import (
	"homework4/internal/dto"
	"homework4/internal/middleware"
	"homework4/internal/service"
	"homework4/pkg/apperror"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

// CreatePost godoc
// @Summary 新增文章
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreatePostRequest true "create post request"
// @Success 201 {object} model.Post
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /posts [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.BadRequest("invalid request body", err))
		return
	}

	userIDValue, _ := c.Get(middleware.ContextUserIDKey)
	userID, ok := userIDValue.(uint)
	if !ok {
		c.Error(apperror.Unauthorized("invalid user context", nil))
		return
	}

	post, err := h.postService.Create(req.Title, req.Content, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, post)
}

// ListPost godoc
// @Summary 获取文章列表
// @Tags Post
// @Produce json
// @Success 200 {array} model.Post
// @Router /posts [get]
func (h *PostHandler) ListPost(c *gin.Context) {
	posts, err := h.postService.List()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost godoc
// @Summary 获取文章详情
// @Tags Post
// @Produce json
// @Param id path int true "post id"
// @Success 200 {object} model.Post
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [get]
func (h *PostHandler) GetPost(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	post, err := h.postService.Detail(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePost godoc
// @Summary 更新文章内容
// @Tags Post
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "post id"
// @Param request body dto.UpdatePostRequest true "update post request"
// @Success 200 {object} model.Post
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [put]
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	var req dto.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.BadRequest("invalid request body", err))
		return
	}

	userIDValue, _ := c.Get(middleware.ContextUserIDKey)
	userID, ok := userIDValue.(uint)
	if !ok {
		c.Error(apperror.Unauthorized("invalid user context", nil))
		return
	}

	post, err := h.postService.Update(id, req.Title, req.Content, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, post)
}

// DeletePost godoc
// @Summary 删除文章
// @Tags Post
// @Produce json
// @Security BearerAuth
// @Param id path int true "post id"
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /posts/{id} [delete]
func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	userIDValue, _ := c.Get(middleware.ContextUserIDKey)
	userID, ok := userIDValue.(uint)
	if !ok {
		c.Error(apperror.Unauthorized("未登录用户", nil))
		return
	}

	if err := h.postService.Delete(id, userID); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func parseID(raw string) (uint, error) {
	id, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, apperror.BadRequest("非法ID", err)
	}
	return uint(id), nil
}
