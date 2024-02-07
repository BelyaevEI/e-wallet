package walletservice_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BelyaevEI/e-wallet/internal/logger"
	"github.com/BelyaevEI/e-wallet/internal/models"
	"github.com/BelyaevEI/e-wallet/internal/storage/database"
	"github.com/BelyaevEI/e-wallet/internal/storage/database/mocks"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletrepository"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTransferFund(t *testing.T) {
	testcase := struct {
		name     string
		wantCode int
	}{
		name:     "Transfer funds",
		wantCode: http.StatusOK,
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	walletstorage := mocks.NewMockWalletStorage(ctrl)
	wallethistory := mocks.NewMockWalletHist(ctrl)
	cache := mocks.NewMockCache(ctrl)

	walletstorage.EXPECT().CheckWallet(ctx, uint32(12345)).Return(false, nil)
	walletstorage.EXPECT().CheckWallet(ctx, uint32(123456)).Return(false, nil)
	walletstorage.EXPECT().CheckAmount(ctx, uint32(12345)).Return(float64(100), nil)
	walletstorage.EXPECT().TransferFunds(ctx, gomock.Any(), gomock.Any()).Return(nil)
	cache.EXPECT().ModifyWallet(gomock.Any(), gomock.Any()).Return(nil)
	cache.EXPECT().ModifyWallet(gomock.Any(), gomock.Any()).Return(nil)

	wallethistory.EXPECT().WriteTransation(ctx, gomock.Any(), gomock.Any()).Return(nil)

	store := &database.Storage{
		WalletStorage: walletstorage,
		WalletHistory: wallethistory,
		Cache:         cache,
	}

	log, _ := logger.New()
	walletRepository := walletrepository.New(store)
	walletservice := walletservice.New(log, walletRepository)

	t.Run(testcase.name, func(t *testing.T) {

		responseRecorder := httptest.NewRecorder()

		r := models.Wallet{ID: uint32(123456), Amount: float64(10)}
		req, _ := json.Marshal(r)
		requestBody := strings.NewReader(string(req))

		request := httptest.NewRequest(http.MethodPost, "/api/v1/wallet/12345/send", requestBody)
		request.Header.Set("Content-Type", "json/application")

		walletservice.TransferFunds(responseRecorder, request)

		result := responseRecorder.Result()
		assert.Equal(t, testcase.wantCode, result.StatusCode) // Check satus code 200
	})

}
