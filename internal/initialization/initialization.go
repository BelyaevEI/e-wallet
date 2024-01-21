package initialization

import (
	"github.com/BelyaevEI/e-wallet/internal/config"
	"github.com/BelyaevEI/e-wallet/internal/logger"
	"github.com/BelyaevEI/e-wallet/internal/route"
	"github.com/BelyaevEI/e-wallet/internal/storage/database"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletrepository"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletservice"
	"github.com/go-chi/chi"
)

type Init struct {
	Host  string
	Port  string
	Route *chi.Mux
}

// Initialization all entites
func GoInit(log *logger.Logger) (Init, error) {

	// Reading config file
	cfg, err := config.LoadConfig("../../")
	if err != nil {
		log.Log.Error("read config file is fail: ", err)
		return Init{}, nil
	}

	// Connect to database
	store, err := database.NewConnect(cfg.DSN)
	if err != nil {
		log.Log.Error("connection to database is failed: ", err)
	}

	// Create wallet repository
	walletRepository := walletrepository.New(store)
	walletservice := walletservice.New(log, walletRepository)

	// Create new router
	route := route.New(walletservice)

	return Init{
		Route: route,
		Host:  cfg.Host,
		Port:  cfg.Port,
	}, nil
}
