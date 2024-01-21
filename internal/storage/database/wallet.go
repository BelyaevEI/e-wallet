package database

import "database/sql"

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
