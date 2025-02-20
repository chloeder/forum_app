package posts

import (
	"net/http"
	"strconv"

	"github.com/chloeder/forum_app/internal/models/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCommentsByPostId(c *gin.Context) {
	ctx := c.Request.Context()

	// Get post id from path
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

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

	// Get comments from service
	comments, err := h.postService.GetCommentsByPostId(ctx, int64(postID), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": comments})
}

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	// Get user id from context
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}

	// Get post id from path
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	// Bind request to struct
	var req posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create comment
	err = h.postService.CreateComment(ctx, userID.(int64), int64(postID), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
