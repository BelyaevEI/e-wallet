package walletservice_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BelyaevEI/e-wallet/internal/logger"
	"github.com/BelyaevEI/e-wallet/internal/models"
	"github.com/BelyaevEI/e-wallet/internal/storage/database"
	"github.com/BelyaevEI/e-wallet/internal/storage/database/mocks"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletrepository"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetWalletHistory(t *testing.T) {
	testcase := struct {
		name     string
		wantCode int
	}{
		name:     "History operation of wallet",
		wantCode: http.StatusOK,
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	walletstorage := mocks.NewMockWalletStorage(ctrl)
	wallethistory := mocks.NewMockWalletHist(ctrl)
	cache := mocks.NewMockCache(ctrl)

	walletstorage.EXPECT().CheckWallet(ctx, uint32(12345)).Return(false, nil)
	wallethistory.EXPECT().Get(ctx, uint32(12345)).Return([]models.WalletHistory{
		{Time: time.Now(),
			FromID: uint32(12345),
			ToID:   uint32(12346),
			Amount: float32(10)}}, nil)

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
		request := httptest.NewRequest(http.MethodGet, "/api/v1/wallet/12345/history", nil)

		walletservice.GetWalletHistory(responseRecorder, request)

		result := responseRecorder.Result()
		assert.Equal(t, testcase.wantCode, result.StatusCode) // Check satus code 200
	})

}
