package walletservice_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BelyaevEI/e-wallet/internal/logger"
	"github.com/BelyaevEI/e-wallet/internal/models"
	"github.com/BelyaevEI/e-wallet/internal/storage/database"
	"github.com/BelyaevEI/e-wallet/internal/storage/database/mocks"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletrepository"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateWallet(t *testing.T) {
	testcase := struct {
		name     string
		wantCode int
	}{
		name:     "Create new wallet",
		wantCode: http.StatusOK,
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	walletstorage := mocks.NewMockWalletStorage(ctrl)
	wallethistory := mocks.NewMockWalletHist(ctrl)
	cache := mocks.NewMockCache(ctrl)

	walletstorage.EXPECT().CheckWallet(ctx, gomock.Any()).Return(true, nil)
	walletstorage.EXPECT().CreateWallet(ctx, gomock.Any()).Return(nil)
	cache.EXPECT().AddWallet(gomock.Any()).Return(nil)

	store := &database.Storage{
		WalletStorage: walletstorage,
		WalletHistory: wallethistory,
		Cache:         cache,
	}

	log, _ := logger.New()
	walletRepository := walletrepository.New(store)
	walletservice := walletservice.New(log, walletRepository)

	t.Run(testcase.name, func(t *testing.T) {
		var res models.Wallet

		request := httptest.NewRequest(http.MethodPost, "/api/v1/wallet", nil)
		responseRecorder := httptest.NewRecorder()

		walletservice.CreateWallet(responseRecorder, request)

		result := responseRecorder.Result()
		resBody, err := io.ReadAll(result.Body)
		defer result.Body.Close()

		json.Unmarshal(resBody, &res)

		// Check for empty body and error
		assert.NotEmpty(t, res)
		require.NoError(t, err)
		assert.Equal(t, testcase.wantCode, result.StatusCode) // Check satus code 200
	})

}
