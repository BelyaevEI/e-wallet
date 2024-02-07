package walletservice

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (service *Service) GetWalletBalance(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	// Get wallet id
	// walletID, err := strconv.Atoi(chi.URLParam(request, "walletid"))
	path := request.URL.Path
	parts := strings.Split(path, "/")
	walletID, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		service.log.Log.Error("reading wallet id from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check correct wallet
	ok, err := service.walletrepository.CheckWallet(ctx, uint32(walletID))
	if err != nil {
		service.log.Log.Error("reading wallet id is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Wallet with this ID not exists
	if ok {
		service.log.Log.Error("wallet with this id not exists")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	//Given current balance of wallet
	balance, err := service.walletrepository.GetBalance(ctx, uint32(walletID))
	if err != nil {
		service.log.Log.Error("reading balance is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// Marshal server response
	enc := json.NewEncoder(writer)
	if err := enc.Encode(balance); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		service.log.Log.Error("serialization is failed: ", err)
		return
	}

}
