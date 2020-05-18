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
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return user, err
}

/*
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}



func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

*/
