package controllers

import (
	"net/http"
	"text/template"
)

// RootHandler : showing login page
func (server *Server) RootHandler(res http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("server/fe/layout.html", "server/fe/login.html")
	if err != nil {
		panic("Unable to run html template " + err.Error())
	}
	tmpl.ExecuteTemplate(res, "layout", server.ProviderIndex)
}
