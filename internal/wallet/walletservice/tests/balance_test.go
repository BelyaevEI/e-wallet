package walletservice_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BelyaevEI/e-wallet/internal/logger"
	"github.com/BelyaevEI/e-wallet/internal/storage/database"
	"github.com/BelyaevEI/e-wallet/internal/storage/database/mocks"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletrepository"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	testcase := struct {
		name     string
		wantCode int
	}{
		name:     "Given balance of wallet",
		wantCode: http.StatusOK,
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	walletstorage := mocks.NewMockWalletStorage(ctrl)
	wallethistory := mocks.NewMockWalletHist(ctrl)
	cache := mocks.NewMockCache(ctrl)

	walletstorage.EXPECT().CheckWallet(ctx, uint32(12345)).Return(false, nil)
	cache.EXPECT().GetBalance(uint32(12345)).Return(float64(100), nil)

	store := &database.Storage{
		WalletStorage: walletstorage,
		WalletHistory: wallethistory,
		Cache:         cache,
	}

	log, _ := logger.New()
	walletRepository := walletrepository.New(store)
	walletservice := walletservice.New(log, walletRepository)

	t.Run(testcase.name, func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "/api/v1/wallet/12345", nil)
		responseRecorder := httptest.NewRecorder()

		walletservice.GetWalletBalance(responseRecorder, request)

		result := responseRecorder.Result()

		assert.Equal(t, testcase.wantCode, result.StatusCode)
	})
}
