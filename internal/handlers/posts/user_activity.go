package posts

import (
	"net/http"
	"strconv"

	"github.com/chloeder/forum_app/internal/models/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) LikedPost(c *gin.Context) {
	ctx := c.Request.Context()

	// Check user_id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}

	// Check post_id
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	// Bind request
	var req posts.UserActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.postService.LikedPost(ctx, int64(postID), userID.(int64), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}
