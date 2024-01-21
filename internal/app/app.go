package app

import (
	"net/http"

	"github.com/BelyaevEI/e-wallet/internal/initialization"
	"github.com/BelyaevEI/e-wallet/internal/logger"
)

func NewApp() (*http.Server, error) {
	// Create new connect to logger
	log, err := logger.New()
	if err != nil {
		return nil, err
	}

	// Initialization additional entites
	init, err := initialization.GoInit(log)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    init.Host + ":" + init.Port,
		Handler: init.Route,
	}
	return server, nil
}
