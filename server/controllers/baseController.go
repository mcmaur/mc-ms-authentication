package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mcmaur/mc-ms-authentication/server/models"
)

// Server : variable for common use
type Server struct {
	DB            *gorm.DB
	Router        *mux.Router
	ProviderIndex *ProviderIndex
	Config        models.Config
}

// ProviderIndex : mind your own business
type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}
