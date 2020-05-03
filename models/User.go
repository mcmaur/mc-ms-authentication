package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User : user info to database
type User struct {
	gorm.Model
	Provider          string
	Email             string `gorm:"type:varchar(100);unique_index"`
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}
