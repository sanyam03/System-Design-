package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"slack-clone/internal/db"
)

type Channel struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	SenderID  uint      `json:"sender_id"`
	ChannelID uint      `json:"channel_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *Message) ChannelIDToString() string {
	return strconv.Itoa(int(m.ChannelID))
}

func CreateChannel(c *gin.Context) {
	var channel Channel
	if err := c.ShouldBindJSON(&channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&channel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create channel"})
		return
	}
	c.JSON(http.StatusCreated, channel)
}

func GetChannels(c *gin.Context) {
	var channels []Channel
	if err := db.DB.Find(&channels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch channels"})
		return
	}
	c.JSON(http.StatusOK, channels)
}


