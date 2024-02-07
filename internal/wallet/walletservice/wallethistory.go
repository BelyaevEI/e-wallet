package walletservice

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/BelyaevEI/e-wallet/internal/models"
)

func (service *Service) GetWalletHistory(writer http.ResponseWriter, request *http.Request) {

	var walletFrom models.Wallet

	ctx := request.Context()

	// Get wallet id
	// walletFromId, err := strconv.Atoi(chi.URLParam(request, "walletid"))
	path := request.URL.Path
	parts := strings.Split(path, "/")
	walletFromId, err := strconv.Atoi(parts[len(parts)-2])
	if err != nil {
		service.log.Log.Error("reading wallet id from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	walletFrom.ID = uint32(walletFromId)

	// Check correct wallet
	ok, err := service.walletrepository.CheckWallet(ctx, walletFrom.ID)
	if err != nil {
		service.log.Log.Error("reading wallet id is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// If wallet not exists
	if ok {
		service.log.Log.Error("wallet with this id not exists")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Getting wallet history from wallet id
	walletHistory, err := service.walletrepository.GetWalletHistory(ctx, walletFrom.ID)
	if err != nil {
		service.log.Log.Error("getting wallet history is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// If history not exists
	if len(walletHistory) == 0 {
		service.log.Log.Info("wallet history not exists")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// Marshal server response
	enc := json.NewEncoder(writer)
	if err := enc.Encode(walletHistory); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		service.log.Log.Error("serialization is failed: ", err)
		return
	}
}
