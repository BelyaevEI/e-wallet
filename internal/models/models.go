package models

import "time"

// Application constants
const (
	ConfigName string = "app"
	ConfigType string = "env"
)

type (
	// Struct of wallet
	Wallet struct {
		ID     uint32  `json:"id"`
		Amount float64 `json:"amount"`
	}

	// Struct of wallet history
	WalletHistory struct {
		Time   time.Time `json:"times"`
		FromID uint32    `json:"from_id"`
		ToID   uint32    `json:"to_id"`
		Amount float32   `json:"amount"`
	}
)
