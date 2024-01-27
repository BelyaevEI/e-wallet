package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BelyaevEI/e-wallet/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type WalletStorage struct {
	db *sql.DB
}

func connect2Wallet(dsn string) (*WalletStorage, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Create table for wallet
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS wallet
					(wallet_id bigint NOT NULL,
				    amount DECIMAL(10, 2) NOT NULL)`)
	if err != nil {
		return nil, err
	}
	return &WalletStorage{
		db: db,
	}, nil
}

func (wallet *WalletStorage) CheckWallet(ctx context.Context, id uint32) (bool, error) {
	// true - wallet not exists
	// false - wallet exists

	var idEx uint32

	row := wallet.db.QueryRowContext(ctx, "SELECT wallet_id FROM wallet WHERE wallet_id = $1", id)
	if err := row.Scan(&idEx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (wallet *WalletStorage) CreateWallet(ctx context.Context, id uint32) error {
	_, err := wallet.db.ExecContext(ctx, "INSERT INTO wallet(wallet_id, amount) VALUES ($1, 100.00)", id)
	return err
}

func (wallet *WalletStorage) CheckAmount(ctx context.Context, id uint32) (float64, error) {

	var amount float64

	row := wallet.db.QueryRowContext(ctx, "SELECT amount FROM wallet WHERE wallet_id = $1", id)
	if err := row.Scan(&amount); err != nil {
		return 0, err
	}
	return amount, nil

}

func (wallet *WalletStorage) TransferFunds(ctx context.Context, walletFrom, walletTo models.Wallet) error {

	// Begin transaction for transfer funds between wallet
	tx, err := wallet.db.Begin()
	if err != nil {
		return err
	}

	// Subtraction amount from wallet
	_, err = tx.ExecContext(ctx, "UPDATE wallet SET amount = amount - $1  WHERE wallet_id = $2", walletTo.Amount, walletFrom.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Add amount to wallet
	_, err = tx.ExecContext(ctx, "UPDATE wallet SET amount = amount + $1  WHERE wallet_id = $2", walletTo.Amount, walletTo.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (wallet *WalletStorage) GetBalance(ctx context.Context, id uint32) (models.Wallet, error) {

	var wal models.Wallet

	row := wallet.db.QueryRowContext(ctx, "SELECT wallet_id, amount FROM wallet WHERE wallet_id = $1", id)
	if err := row.Scan(&wal.ID, &wal.Amount); err != nil {
		return models.Wallet{}, err
	}
	return wal, nil
}
