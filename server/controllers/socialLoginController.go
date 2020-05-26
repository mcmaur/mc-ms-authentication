package controllers

import (
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/mcmaur/mc-ms-authentication/server/models"
)

// SocialredirectHandler : redirect to social login page of the provider chosen
func (server *Server) SocialredirectHandler(res http.ResponseWriter, req *http.Request) {
	if _, err := gothic.CompleteUserAuth(res, req); err == nil {
		res.Header().Set("Location", "/user_profile")
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		log.Println("ERR: ", err)
		gothic.BeginAuthHandler(res, req)
	}
}

// SocialCallbackHandler : function executed after return from social network
func (server *Server) SocialCallbackHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Println(res, err)
		return
	}

	var currentUser models.User
	currentUser.FromGothUser(user)
	server.DB.FirstOrCreate(&currentUser)

	err = models.CreateToken(res, req, currentUser.ID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Location", "/user_profile")
	res.WriteHeader(http.StatusTemporaryRedirect)
}
