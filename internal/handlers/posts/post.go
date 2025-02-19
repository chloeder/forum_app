package posts

import (
	"net/http"
	"strconv"

	"github.com/chloeder/forum_app/internal/models/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPosts(c *gin.Context) {
	ctx := c.Request.Context()

	// Default Value
	limit, offset := 10, 0

	// Parse limit from query params
	if limitStr := c.Query("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}

		// Validate limit
		if parsedLimit > 100 {
			limit = 100
		} else {
			limit = parsedLimit
		}
	}

	// Parse offset from query params
	if offsetStr := c.Query("offset"); offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
			return
		}

		// Validate offset
		if parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	// Get posts from service
	posts, err := h.postService.GetPosts(ctx, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": posts})
}

func (h *Handler) GetPostById (c *gin.Context) {
	ctx := c.Request.Context()

	// Get post id from path
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	// Get post from service
	post, err := h.postService.GetPostById(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": post})
}

func (h *Handler) CreatePost (c *gin.Context) {
	ctx := c.Request.Context()

	// Get user id from context
	userID, _ := c.Get("user_id")

	// Bind request body to CreatePostRequest
	var req posts.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	err := h.postService.CreatePost(ctx, userID.(int64), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "post created"})
}

func (h *Handler) UpdatePost (c *gin.Context) {
	ctx := c.Request.Context()

	// Get post id from path
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	// Bind request body to PostModel
	var post posts.UpdatePostRequest
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update post
	err = h.postService.UpdatePost(ctx, int64(id), &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post updated"})
}

func (h *Handler) DeletePost (c *gin.Context) {
	ctx := c.Request.Context()

	// Get post id from path
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	// Delete post
	err = h.postService.DeletePost(ctx, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}
