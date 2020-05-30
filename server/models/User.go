package models

import (
	"errors"
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
func (user *User) FromGothUser(gothuser goth.User) error {
	user.Provider = gothuser.Provider
	user.Email = gothuser.Email
	user.Name = gothuser.Name
	user.FirstName = gothuser.FirstName
	user.LastName = gothuser.LastName
	user.NickName = gothuser.NickName
	user.Description = gothuser.Description
	user.UserID = gothuser.UserID
	user.AvatarURL = gothuser.AvatarURL
	user.Location = gothuser.Location
	user.AccessToken = gothuser.AccessToken
	user.AccessTokenSecret = gothuser.AccessTokenSecret
	user.RefreshToken = gothuser.RefreshToken
	user.ExpiresAt = gothuser.ExpiresAt
	user.IDToken = gothuser.IDToken
	return nil
}

// FindUserByID : find the user by id
func (user *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return user, err
}

// FindUserByEMail : find the user by email
func (user *User) FindUserByEMail(db *gorm.DB, email string) (*User, error) {
	err := db.Debug().Model(User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return user, err
}

// DeleteUserByID : find the user by id and delete it
func (user *User) DeleteUserByID(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// FindAllUsers : find the user by id and delete it
func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// SaveUser : create a new user
func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// UpdateUser : update user infos
func (user *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"email":      user.Email,
			"name":       user.Name,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"nick_name":  user.NickName,
			"update_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	err := db.Model(&User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
