package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"slack-clone/internal/db"
)

func GetMessagesByChannel(c *gin.Context) {
	channelID := c.Param("channel_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var messages []db.Message
	offset := (page - 1) * limit

	if err := db.DB.Where("channel_id = ?", channelID).Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func CreateMessage(c *gin.Context) {
	var message db.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the message to the database
	if err := db.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}

	c.JSON(http.StatusCreated, message)
}