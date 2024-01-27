package walletrepository

import (
	"context"
	"errors"
	"sync"

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

	// Create new wallet
	err = repo.store.WalletStorage.CreateWallet(ctx, wallet.ID)
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

	err := repo.store.WalletStorage.TransferFunds(ctx, walletFrom, walletTo)
	if err != nil {
		return err
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

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	return repo.store.WalletStorage.GetBalance(ctx, id)
}
