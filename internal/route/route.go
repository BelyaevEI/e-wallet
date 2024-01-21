package route

import (
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletservice"
	"github.com/go-chi/chi"
)

func New(walletservice *walletservice.Service) *chi.Mux {

	// New router
	route := chi.NewRouter()

	// Handlers

	return route
}
