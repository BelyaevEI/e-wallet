package walletservice

import (
	"encoding/json"
	"net/http"

	"github.com/BelyaevEI/e-wallet/internal/models"
)

// Create new wallet
func (service *Service) CreateWallet(writer http.ResponseWriter, request *http.Request) {

	var (
		// buf    bytes.Buffer
		wallet models.Wallet
	)

	ctx := request.Context()

	// Getting data for create wallet
	// _, err := buf.ReadFrom(request.Body)
	// if err != nil {
	// 	service.log.Log.Error("reading body from request is failed: ", err)
	// 	writer.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// Generating unique id wallet
	wallet.ID = service.walletrepository.GenerateUniqueID()

	// Create wallet
	err := service.walletrepository.CreateWallet(ctx, wallet)
	if err != nil {
		service.log.Log.Error("creating wallet is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	wallet.Amount = 100.00

	// Serializing answer server
	enc := json.NewEncoder(writer)
	if err := enc.Encode(wallet); err != nil {
		service.log.Log.Error("marshal data is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)

}
