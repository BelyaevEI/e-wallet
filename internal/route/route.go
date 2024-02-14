package route

import (
	_ "github.com/BelyaevEI/e-wallet/docs"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletservice"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(walletservice *walletservice.Service) *chi.Mux {

	// New router
	route := chi.NewRouter()

	// Handlers
	route.Post("/api/v1/wallet", walletservice.CreateWallet)                       // Create new wallet
	route.Post("/api/v1/wallet/{walletid}/send", walletservice.TransferFunds)      // Transfer funds
	route.Get("/api/v1/wallet/{walletid}/history", walletservice.GetWalletHistory) // Given history of wallet
	route.Get("/api/v1/wallet/{walletid}", walletservice.GetWalletBalance)         // Given wallet balance
	route.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))
	return route
}
