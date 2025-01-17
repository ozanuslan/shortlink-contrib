// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wallet_di

import (
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

// Injectors from wire.go:

func InitializeWalletService() (*WalletService, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	walletService, err := NewWalletService(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return walletService, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type WalletService struct {
	Log logger.Logger
}

// WalletService =======================================================================================================
var WalletSet = wire.NewSet(di.DefaultSet, NewWalletService)

func NewWalletService(log logger.Logger) (*WalletService, error) {
	log.Info("Start wallet service")

	return &WalletService{
		Log: log,
	}, nil
}
