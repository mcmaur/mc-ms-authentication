package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mcmaur/mc-ms-authentication/server/models"
	"github.com/mcmaur/mc-ms-authentication/server/response"
)

// GetUserByIDHandler : select user by id and return detailed infos
func (server *Server) GetUserByIDHandler(res http.ResponseWriter, req *http.Request) {
	useridString := mux.Vars(req)["userid"]

	userid, err := strconv.ParseUint(useridString, 10, 32)
	if err != nil {
		log.Println("Error converting string userid to int: " + err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{}
	foundUser, err := user.FindUserByID(server.DB, uint32(userid))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	response.JSONResp(res, http.StatusOK, foundUser)
}

// GetUserByEmailHandler : select user by email addresses and return detailed infos
func (server *Server) GetUserByEmailHandler(res http.ResponseWriter, req *http.Request) {
	userEmail := mux.Vars(req)["useremail"]

	user := models.User{}
	foundUser, err := user.FindUserByEMail(server.DB, userEmail)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	response.JSONResp(res, http.StatusOK, foundUser)
}

// DeleteUserByIDHandler : delete user filtering by user id
func (server *Server) DeleteUserByIDHandler(res http.ResponseWriter, req *http.Request) {
	useridString := mux.Vars(req)["userid"]

	userid, err := strconv.ParseUint(useridString, 10, 32)
	if err != nil {
		log.Println("Error converting string userid to int: " + err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{}
	foundUser, err := user.DeleteUserByID(server.DB, uint32(userid))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	response.JSONResp(res, http.StatusOK, foundUser)
}
