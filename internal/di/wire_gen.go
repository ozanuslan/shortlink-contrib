// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"context"
	"fmt"
	"github.com/batazor/shortlink/internal/api/infrastructure/store"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/metadata/infrastructure/store"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/batazor/shortlink/pkg/rpc"
	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
	"net/http"
	"net/http/pprof"
	"time"
)

// Injectors from default.go:

func InitializeFullService() (*Service, func(), error) {
	context, cleanup, err := NewContext()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := InitLogger(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := InitMQ(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := InitSentry()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := InitMonitoring(handler)
	tracer, cleanup5, err := traicing.NewTracer(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	store, cleanup6, err := InitStore(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := InitProfiling()
	diDiAutoMaxPro, cleanup7, err := InitAutoMaxProcs(logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup8, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup9, err := rpc.InitClient(logger, tracer)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewFullService(context, logger, mq, handler, serveMux, tracer, store, pprofEndpoint, diDiAutoMaxPro, rpcServer, clientConn)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_api.go:

func InitializeAPIService() (*Service, func(), error) {
	context, cleanup, err := NewContext()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := InitLogger(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := InitMQ(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := InitSentry()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := InitMonitoring(handler)
	tracer, cleanup5, err := traicing.NewTracer(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	store, cleanup6, err := InitStore(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkStore, err := InitLinkStore(context, logger, store)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := InitProfiling()
	diDiAutoMaxPro, cleanup7, err := InitAutoMaxProcs(logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup8, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup9, err := rpc.InitClient(logger, tracer)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewAPIService(context, logger, mq, handler, serveMux, tracer, store, linkStore, pprofEndpoint, diDiAutoMaxPro, rpcServer, clientConn)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_bot.go:

func InitializeBotService() (*Service, func(), error) {
	context, cleanup, err := NewContext()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := InitLogger(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := InitMQ(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	diDiAutoMaxPro, cleanup4, err := InitAutoMaxProcs(logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewBotService(logger, mq, diDiAutoMaxPro)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_logger.go:

func InitializeLoggerService() (*Service, func(), error) {
	context, cleanup, err := NewContext()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := InitLogger(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := InitMQ(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	diDiAutoMaxPro, cleanup4, err := InitAutoMaxProcs(logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewLoggerService(logger, mq, diDiAutoMaxPro)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_metadata.go:

func InitializeMetadataService() (*Service, func(), error) {
	context, cleanup, err := NewContext()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := InitLogger(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	diDiAutoMaxPro, cleanup3, err := InitAutoMaxProcs(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	store, cleanup4, err := InitStore(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup5, err := traicing.NewTracer(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup6, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metaStore, err := InitMetaStore(context, logger, store)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup7, err := InitSentry()
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := InitMonitoring(handler)
	service, err := NewMetadataService(logger, diDiAutoMaxPro, store, rpcServer, metaStore, serveMux, handler)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// default.go:

// Service - heplers
type Service struct {
	Ctx    context.Context
	Log    logger.Logger
	Tracer *opentracing.Tracer
	// TracerClose func()
	Sentry        *sentryhttp.Handler
	DB            *db.Store
	LinkStore     *store.LinkStore
	MetaStore     *meta_store.MetaStore
	MQ            mq.MQ
	ServerRPC     *rpc.RPCServer
	ClientRPC     *grpc.ClientConn
	Monitoring    *http.ServeMux
	PprofEndpoint PprofEndpoint
}

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// Cobra/Flags =========================================================================================================
func InitFlags() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	return rootCmd, nil
}

// Monitoring ==========================================================================================================
func InitMonitoring(sentryHandler *sentryhttp.Handler) *http.ServeMux {

	registry := prometheus.NewRegistry()

	health := healthcheck.NewMetricsHandler(registry, "common")

	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	commonMux := http.NewServeMux()

	commonMux.Handle("/metrics", sentryHandler.Handle(promhttp.Handler()))

	commonMux.HandleFunc("/live", sentryHandler.HandleFunc(health.LiveEndpoint))

	commonMux.HandleFunc("/ready", sentryHandler.HandleFunc(health.ReadyEndpoint))

	return commonMux
}

// Profiling ===========================================================================================================
type PprofEndpoint *http.ServeMux

func InitProfiling() PprofEndpoint {

	pprofMux := http.NewServeMux()

	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return pprofMux
}

// Health ==============================================================================================================
func NewHealthCheck() (healthcheck.Handler, error) {

	endpoint := healthcheck.NewHandler()

	go http.ListenAndServe("0.0.0.0:9090", endpoint)

	return endpoint, nil
}

// AutoMaxProcs ========================================================================================================
type diAutoMaxPro *string

// InitAutoMaxProcs - Automatically set GOMAXPROCS to match Linux container CPU quota
func InitAutoMaxProcs(log logger.Logger) (diAutoMaxPro, func(), error) {
	undo, err := maxprocs.Set(maxprocs.Logger(func(s string, args ...interface{}) {
		log.Info(fmt.Sprintf(s, args))
	}))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		undo()
	}

	return nil, cleanup, nil
}

// InitLinkStore
func InitLinkStore(ctx context.Context, log logger.Logger, conn *db.Store) (*store.LinkStore, error) {
	st := store.LinkStore{}
	linkStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

// InitMetaStore
func InitMetaStore(ctx context.Context, log logger.Logger, conn *db.Store) (*meta_store.MetaStore, error) {
	st := meta_store.MetaStore{}
	metaStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return metaStore, nil
}

// Logger ==============================================================================================================
func InitLogger(ctx context.Context) (logger.Logger, func(), error) {
	viper.SetDefault("LOG_LEVEL", logger.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := logger.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {

		_ = log.Close()
	}

	return log, cleanup, nil
}

func InitSentry() (*sentryhttp.Handler, func(), error) {
	viper.SetDefault("SENTRY_DSN", "")
	DSN := viper.GetString("SENTRY_DSN")

	if DSN == "" {
		return nil, func() {}, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: DSN,
	})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		sentry.Flush(time.Second * 5)
		sentry.Recover()
	}

	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	return sentryHandler, cleanup, nil
}

// Store ===============================================================================================================
// InitStore return db
func InitStore(ctx context.Context, log logger.Logger) (*db.Store, func(), error) {
	var st db.Store
	db2, err := st.Use(ctx, log)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db2.Store.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db2, cleanup, nil
}

// MQ ==================================================================================================================
func InitMQ(ctx context.Context, log logger.Logger) (mq.MQ, func(), error) {
	viper.SetDefault("MQ_ENABLED", "false")

	if viper.GetBool("MQ_ENABLED") {
		var service mq.DataBus
		dataBus, err := service.Use(ctx, log)
		if err != nil {
			return nil, func() {}, err
		}

		cleanup := func() {
			if err := dataBus.Close(); err != nil {
				log.Error(err.Error())
			}
		}

		return dataBus, cleanup, nil
	}

	return nil, func() {}, nil
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(NewContext, InitAutoMaxProcs, InitLogger, traicing.NewTracer)

// FullService =========================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService,
	InitStore,
	InitSentry,
	InitMonitoring,
	InitProfiling,
	InitMQ, rpc.InitServer, rpc.InitClient,
)

func NewFullService(
	ctx context.Context,
	log logger.Logger, mq2 mq.MQ,

	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer, db2 *db.Store,

	pprofHTTP PprofEndpoint,
	autoMaxProcsOption diAutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:    ctx,
		Log:    log,
		MQ:     mq2,
		Tracer: tracer,

		Monitoring: monitoring,
		Sentry:     sentryHandler,
		DB:         db2,

		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

// service_api.go:

// APIService =======================================================================================================
var APISet = wire.NewSet(
	DefaultSet,
	InitStore,
	InitLinkStore,
	InitSentry,
	InitMonitoring,
	InitProfiling,
	InitMQ, rpc.InitServer, rpc.InitClient, NewAPIService,
)

func NewAPIService(
	ctx context.Context,
	log logger.Logger, mq2 mq.MQ,

	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer, db2 *db.Store,
	linkStore *store.LinkStore,
	pprofHTTP PprofEndpoint,
	autoMaxProcsOption diAutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:    ctx,
		Log:    log,
		MQ:     mq2,
		Tracer: tracer,

		Monitoring:    monitoring,
		Sentry:        sentryHandler,
		DB:            db2,
		LinkStore:     linkStore,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

// service_bot.go:

// BotService ==========================================================================================================
var BotSet = wire.NewSet(DefaultSet, NewBotService, InitMQ)

func NewBotService(log logger.Logger, mq2 mq.MQ, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq2,
	}, nil
}

// service_logger.go:

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(DefaultSet, InitMQ, NewLoggerService)

func NewLoggerService(log logger.Logger, mq2 mq.MQ, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq2,
	}, nil
}

// service_metadata.go:

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(
	DefaultSet,
	InitStore, rpc.InitServer, InitMetaStore,
	InitSentry,
	InitMonitoring,
	NewMetadataService,
)

func NewMetadataService(
	log logger.Logger,
	autoMaxProcsOption diAutoMaxPro, db2 *db.Store,
	serverRPC *rpc.RPCServer,
	metaStore *meta_store.MetaStore,
	monitoring *http.ServeMux,
	sentryHandler *sentryhttp.Handler,
) (*Service, error) {
	return &Service{
		Log:        log,
		ServerRPC:  serverRPC,
		DB:         db2,
		MetaStore:  metaStore,
		Monitoring: monitoring,
		Sentry:     sentryHandler,
	}, nil
}
