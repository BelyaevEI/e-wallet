package walletrepository

import (
	"sync"

	"github.com/BelyaevEI/e-wallet/internal/storage/database"
)

type WalletRepository struct {
	store *database.Storage
	mutex sync.RWMutex
}

func New(store *database.Storage) *WalletRepository {
	return &WalletRepository{
		store: store,
	}
}
