package db

import (
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
type Channel struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	CreatedAt time.Time `json:"created_at"`
}

// User struct
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Message struct
type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	SenderID  uint      `json:"sender_id"`   // References a User
	ChannelID uint      `json:"channel_id"` // References a Channel
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
func (m *Message) ChannelIDToString() string {
	return strconv.Itoa(int(m.ChannelID))
}

func InitDB() {
	dsn := "host=localhost user=sanyamdilipkumarbharani password=root dbname=slack_clone port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected!")

	// Automigrate models
	AutoMigrate(&Channel{}, &User{}, &Message{})
}

func AutoMigrate(models ...interface{}) {
	if err := DB.AutoMigrate(models...); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
