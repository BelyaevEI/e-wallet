package database

import "database/sql"

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
					(timestamp_column TIMESTAMP WITH TIME ZONE,
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
