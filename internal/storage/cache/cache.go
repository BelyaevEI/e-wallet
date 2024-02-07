package cache

import (
	"strconv"

	"github.com/BelyaevEI/e-wallet/internal/models"
	"github.com/go-redis/redis"
)

type Cache interface {
	AddWallet(wallet models.Wallet) error
	FillCache(wallets []models.Wallet) error
	GetBalance(walletID uint32) (float64, error)
	ModifyWallet(walletID uint32, amount float64) error
}

type Redis struct {
	rdb *redis.Client
}

func NewConnect(addr, password string) *Redis {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &Redis{rdb: rdb}
}

func (r *Redis) AddWallet(wallet models.Wallet) error {
	return r.rdb.Set(strconv.Itoa(int(wallet.ID)), strconv.FormatFloat(wallet.Amount, 'f', -1, 64), 0).Err()
}

func (r *Redis) ModifyWallet(walletID uint32, amount float64) error {

	// Get current amount of wallet
	stramount, err := r.rdb.Get(strconv.Itoa(int(walletID))).Result()
	if err != nil {
		return err
	}

	am, err := strconv.ParseFloat(stramount, 64)
	if err != nil {
		return err
	}

	// Add amount to current value amount of wallet
	return r.rdb.Set(strconv.Itoa(int(walletID)), strconv.FormatFloat(am+amount, 'f', -1, 64), 0).Err()
}

func (r *Redis) GetBalance(walletID uint32) (float64, error) {

	stramount, err := r.rdb.Get(strconv.Itoa(int(walletID))).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(stramount, 64)
}

func (r *Redis) FillCache(wallets []models.Wallet) error {
	for _, wallet := range wallets {

		err := r.AddWallet(wallet)
		if err != nil {
			return err
		}
	}
	return nil
}
