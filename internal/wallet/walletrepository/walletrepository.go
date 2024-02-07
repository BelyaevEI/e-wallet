package walletrepository

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/BelyaevEI/e-wallet/internal/models"
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

// Create wallet from data request
func (repo *WalletRepository) CreateWallet(ctx context.Context, wallet models.Wallet) error {

	// Check exists wallet
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	ok, err := repo.store.WalletStorage.CheckWallet(ctx, wallet.ID)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("wallet with this id is exists")
	}

	// Create new wallet in database
	err = repo.store.WalletStorage.CreateWallet(ctx, wallet.ID)
	if err != nil {
		return err
	}

	wallet.Amount = 100

	// Add wallet to cache
	err = repo.store.Cache.AddWallet(wallet)
	if err != nil {
		return err
	}
	return nil
}

func (repo *WalletRepository) CheckWallets(ctx context.Context, walletFrom, walletTo models.Wallet) error {

	// Check exists wallet
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	ok, err := repo.store.WalletStorage.CheckWallet(ctx, walletFrom.ID)
	if err != nil {
		return err
	}

	if ok {
		return errors.New("wallet with this id is not exists")
	}

	// Check exists wallets
	ok, err = repo.store.WalletStorage.CheckWallet(ctx, walletTo.ID)
	if err != nil {
		return err
	}

	if ok {
		return errors.New("wallet with this id is not exists")
	}

	// Check needs amount
	amount, err := repo.store.WalletStorage.CheckAmount(ctx, walletFrom.ID)
	if err != nil {
		return err
	}

	if amount-float64(walletFrom.Amount) < 0 {
		return errors.New("insufficient funds in the wallet")
	}

	return nil
}

func (repo *WalletRepository) TransferFunds(ctx context.Context, walletFrom, walletTo models.Wallet) error {

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	// Change amount of wallet in DB
	err := repo.store.WalletStorage.TransferFunds(ctx, walletFrom, walletTo)
	if err != nil {
		return err
	}

	// Change amount of wallet in cache
	// Source wallet
	err = repo.store.Cache.ModifyWallet(walletFrom.ID, -walletTo.Amount)
	if err != nil {
		return nil
	}

	// Receiver wallet
	err = repo.store.Cache.ModifyWallet(walletTo.ID, walletTo.Amount)
	if err != nil {
		return nil
	}

	return nil
}

func (repo *WalletRepository) WriteTransation(ctx context.Context, walletFrom, walletTo models.Wallet) error {

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	err := repo.store.WalletHistory.WriteTransation(ctx, walletFrom, walletTo)
	if err != nil {
		return err
	}

	return nil
}

func (repo *WalletRepository) CheckWallet(ctx context.Context, id uint32) (bool, error) {

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	ok, err := repo.store.WalletStorage.CheckWallet(ctx, id)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (repo *WalletRepository) GetWalletHistory(ctx context.Context, id uint32) ([]models.WalletHistory, error) {

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	history, err := repo.store.WalletHistory.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return history, nil

}

func (repo *WalletRepository) GetBalance(ctx context.Context, id uint32) (models.Wallet, error) {

	var wallet models.Wallet

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	// Given balance from cache
	balance, err := repo.store.Cache.GetBalance(id)
	if err != nil {

		//Given balance from db
		return repo.store.WalletStorage.GetBalance(ctx, id)
	}

	wallet.ID, wallet.Amount = id, balance
	return wallet, nil
}

func (repo *WalletRepository) GenerateUniqueID() uint32 {

	time := time.Now().UnixNano()

	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Convert random number into uint32
	randomNumber := binary.BigEndian.Uint32(randomBytes)

	// Convert to time into uint32 and adding random number
	uniqueNumber := uint32(time) + randomNumber

	return uniqueNumber
}

func (repo *WalletRepository) FillCache() {

	// Getting all wallet from db
	wallets, err := repo.store.WalletStorage.GetAllWallet()
	if err != nil {
		return
	}

	// Fill the cache with data from the database
	err = repo.store.Cache.FillCache(wallets)
	if err != nil {
		return
	}

}
