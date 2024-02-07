package walletservice

import (
	"github.com/BelyaevEI/e-wallet/internal/logger"
	"github.com/BelyaevEI/e-wallet/internal/wallet/walletrepository"
)

type Service struct {
	log              *logger.Logger
	walletrepository *walletrepository.WalletRepository
}

func New(log *logger.Logger, walletrepository *walletrepository.WalletRepository) *Service {

	return &Service{
		log:              log,
		walletrepository: walletrepository,
	}
}
