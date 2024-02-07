package database

import (
	"github.com/BelyaevEI/e-wallet/internal/config"
	"github.com/BelyaevEI/e-wallet/internal/storage/cache"
)

type Storage struct {
	WalletStorage WalletStorage
	WalletHistory WalletHist
	Cache         cache.Cache
}

// Connect to postgresql database
func NewConnect(cfg config.Config) (*Storage, error) {

	//Connect to database wallet
	wallet, err := connect2Wallet(cfg.DSN)
	if err != nil {
		return &Storage{}, nil
	}

	// Connect to database wallet history
	history, err := connect2WalletHistory(cfg.DSN)
	if err != nil {
		return &Storage{}, nil
	}

	// Connect to cache redis
	rdb := cache.NewConnect(cfg.RedisAddr, cfg.RedisPassword)

	return &Storage{
		WalletStorage: wallet,
		WalletHistory: history,
		Cache:         rdb,
	}, nil
}
