package controllers

import (
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/mcmaur/mc-ms-authentication/server/models"
)

// LogoutHandler : logut function
func (server *Server) LogoutHandler(res http.ResponseWriter, req *http.Request) {
	gothic.Logout(res, req)
	models.DeleteToken(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}
