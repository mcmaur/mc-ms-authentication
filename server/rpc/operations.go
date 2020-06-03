package rpc

import (
	"github.com/mcmaur/mc-ms-authentication/server/models"
)

// GetUserByIDRpc : select user by id and return detailed infos
func (server *Server) GetUserByIDRpc(userid uint32, user *models.User) error {
	newuser := models.User{}
	user, err := newuser.FindUserByID(server.DB, userid)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmailRpc : select user by email addresses and return detailed infos
func (server *Server) GetUserByEmailRpc(userEmail string, user *models.User) error {
	newuser := models.User{}
	user, err := newuser.FindUserByEMail(server.DB, userEmail)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserByIDRpc : delete user filtering by user id
func (server *Server) DeleteUserByIDRpc(userid uint32, rows int) error {
	newuser := models.User{}
	user, err := newuser.DeleteUserByID(server.DB, uint32(userid))
	if err != nil {
		return err
	}
	return nil
}
