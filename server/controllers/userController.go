package controllers

import (
	"net/http"
	"text/template"

	"github.com/mcmaur/mc-ms-authentication/server/models"
)

// UserProfileHandler : showing page with user infos
func (server *Server) UserProfileHandler(res http.ResponseWriter, req *http.Request) {

	userid, err := models.ExtractTokenID(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{}
	foundUser, err := user.FindUserByID(server.DB, userid)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("server/fe/layout.html", "server/fe/user_info.html")
	if err != nil {
		panic("Unable to run html template " + err.Error())
	}
	tmpl.ExecuteTemplate(res, "layout", foundUser)
}
