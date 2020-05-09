package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/markbates/goth"
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

// FromGothUser : trasforming from goth user to this model
func (currentUser *User) FromGothUser(user goth.User) error {
	currentUser.Provider = user.Provider
	currentUser.Email = user.Email
	currentUser.Name = user.Name
	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.NickName = user.NickName
	currentUser.Description = user.Description
	currentUser.UserID = user.UserID
	currentUser.AvatarURL = user.AvatarURL
	currentUser.Location = user.Location
	currentUser.AccessToken = user.AccessToken
	currentUser.AccessTokenSecret = user.AccessTokenSecret
	currentUser.RefreshToken = user.RefreshToken
	currentUser.ExpiresAt = user.ExpiresAt
	currentUser.IDToken = user.IDToken
}
