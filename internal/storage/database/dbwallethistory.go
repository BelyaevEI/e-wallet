package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BelyaevEI/e-wallet/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type WalletHist interface {
	Get(ctx context.Context, id uint32) ([]models.WalletHistory, error)
	WriteTransation(ctx context.Context, walletFrom models.Wallet, walletTo models.Wallet) error
}

type WalletHistory struct {
	db *sql.DB
}

func connect2WalletHistory(dsn string) (*WalletHistory, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Create table for wallet history
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS history
					(times TIMESTAMP WITH TIME ZONE,
				    from_id bigint NOT NULL,
					to_id bigint NOT NULL, 
					amount DECIMAL(10, 2) NOT NULL)`)
	if err != nil {
		return nil, err
	}
	return &WalletHistory{
		db: db,
	}, nil
}

func (history *WalletHistory) WriteTransation(ctx context.Context, walletFrom, walletTo models.Wallet) error {

	_, err := history.db.ExecContext(ctx, `INSERT INTO history(times, from_id, to_id, amount) 
											VALUES (CURRENT_TIMESTAMP, $1, $2, $3)`,
		walletFrom.ID, walletTo.ID, walletTo.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (history *WalletHistory) Get(ctx context.Context, id uint32) ([]models.WalletHistory, error) {

	historyWallet := make([]models.WalletHistory, 0)

	rows, err := history.db.QueryContext(ctx, `SELECT times, from_id, to_id, amount FROM history 
											WHERE from_id = $1 OR to_id = $1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var historyWal models.WalletHistory
		err := rows.Scan(&historyWal.Time, &historyWal.FromID, &historyWal.ToID, &historyWal.Amount)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, err
			}
			return nil, nil
		}

		historyWallet = append(historyWallet, historyWal)
	}

	return historyWallet, nil
}
