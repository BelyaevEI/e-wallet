package database

type Store interface {
}

type Storage struct {
	WalletStorage WalletStorage
	WalletHistory WalletHistory
}

// Connect to postgresql database
func NewConnect(dsn string) (*Storage, error) {

	//Connect to database wallet
	wallet, err := connect2Wallet(dsn)
	if err != nil {
		return &Storage{}, nil
	}

	// Connect to database wallet history
	history, err := connect2WalletHistory(dsn)
	if err != nil {
		return &Storage{}, nil
	}

	return &Storage{
		WalletStorage: *wallet,
		WalletHistory: *history,
	}, nil
}
