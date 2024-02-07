package walletservice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/BelyaevEI/e-wallet/internal/models"
)

func (service *Service) TransferFunds(writer http.ResponseWriter, request *http.Request) {

	var (
		buf        bytes.Buffer
		walletTo   models.Wallet
		walletFrom models.Wallet
	)

	ctx := request.Context()

	// Get wallet id from transfer funds
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

	// Getting data for transfer funds
	_, err = buf.ReadFrom(request.Body)
	if err != nil {
		service.log.Log.Error("reading body from request is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshal data
	if err := json.Unmarshal(buf.Bytes(), &walletTo); err != nil {
		service.log.Log.Error("unmarshal data is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check wallets exists and balance is pozitiv
	err = service.walletrepository.CheckWallets(ctx, walletFrom, walletTo)
	if err != nil {
		service.log.Log.Error("checking wallets or amount is failed: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Transfer funds
	err = service.walletrepository.TransferFunds(ctx, walletFrom, walletTo)
	if err != nil {
		service.log.Log.Error("transfer funds is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write transaction to history
	err = service.walletrepository.WriteTransation(ctx, walletFrom, walletTo)
	if err != nil {
		service.log.Log.Error("write transaction is failed: ", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Success transfer funds to wallet
	writer.WriteHeader(http.StatusOK)
}
