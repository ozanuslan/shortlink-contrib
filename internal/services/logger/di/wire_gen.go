// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package logger_di

import (
	"context"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/v1"
	"github.com/shortlink-org/shortlink/internal/services/logger/application"
	"github.com/shortlink-org/shortlink/internal/services/logger/infrastructure/mq"
)

// Injectors from wire.go:

func InitializeLoggerService() (*LoggerService, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	service, err := NewLoggerApplication(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := mq_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	event, err := InitLoggerMQ(context, logger, mq, service)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	loggerService, err := NewLoggerService(logger, service, event)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return loggerService, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type LoggerService struct {
	// Common
	Log logger.Logger

	// Delivery
	loggerMQ *logger_mq.Event

	// Application
	loggerService *logger_application.Service
}

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(di.DefaultSet, mq_di.New, InitLoggerMQ,

	NewLoggerApplication,

	NewLoggerService,
)

func InitLoggerMQ(ctx2 context.Context, log logger.Logger, mq v1.MQ, service *logger_application.Service) (*logger_mq.Event, error) {
	loggerMQ, err := logger_mq.New(mq, log, service)
	if err != nil {
		return nil, err
	}

	return loggerMQ, nil
}

func NewLoggerApplication(logger2 logger.Logger) (*logger_application.Service, error) {
	loggerService, err := logger_application.New(logger2)
	if err != nil {
		return nil, err
	}

	return loggerService, nil
}

func NewLoggerService(

	log logger.Logger,

	loggerService *logger_application.Service,

	loggerMQ *logger_mq.Event,
) (*LoggerService, error) {
	return &LoggerService{

		Log: log,

		loggerService: loggerService,

		loggerMQ: loggerMQ,
	}, nil
}
