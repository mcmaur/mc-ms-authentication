package rpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/mcmaur/mc-ms-authentication/server/models"
)

func (server *Server) GetLine(id uint32, returnUser *models.User) error {
	fmt.Printf("Receive: %v\n", id)

	user := models.User{}
	returnUser, err := user.FindUserByID(server.DB, id)
	if err != nil {
		//http.Error(res, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

/*
func startold() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	rpc.Register(server)
	rpc.Accept(inbound)
}*/

func start() {
	user := new(models.User)
	// Publish the receivers methods
	err := rpc.Register(user)
	if err != nil {
		log.Fatal("Format of service User isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	log.Printf("Serving RPC server on port %d", 1234)
	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
