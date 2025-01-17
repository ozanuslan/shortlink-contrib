//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Link Service DI-package
*/
package link_di

import (
	"context"

	"github.com/authzed/authzed-go/v1"
	"github.com/go-redis/cache/v9"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	mq_di "github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/store"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"
	"github.com/shortlink-org/shortlink/internal/services/link/application/link"
	"github.com/shortlink-org/shortlink/internal/services/link/application/link_cqrs"
	"github.com/shortlink-org/shortlink/internal/services/link/application/sitemap"
	api_mq "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/mq"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/cqrs/query"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud"
	cqrs "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/run"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	metadata_rpc "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
)

type LinkService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Security
	authPermission *authzed.Client

	// Delivery
	linkMQ            *api_mq.Event
	run               *run.Response
	linkRPCServer     *link_rpc.Link
	linkCQRSRPCServer *cqrs.Link
	sitemapRPCServer  *sitemap_rpc.Sitemap

	// Application
	linkService     *link.Service
	linkCQRSService *link_cqrs.Service
	sitemapService  *sitemap.Service

	// Repository
	linkStore crud.Repository

	// CQRS
	cqsStore   *cqs.Store
	queryStore *query.Store
}

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(
	di.DefaultSet,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
	store.New,

	// Delivery
	InitLinkMQ,

	NewLinkRPCServer,
	NewLinkCQRSRPCServer,
	NewSitemapRPCServer,
	NewRunRPCServer,

	NewLinkRPCClient,
	NewMetadataRPCClient,

	// Applications
	NewLinkApplication,
	NewLinkCQRSApplication,
	NewSitemapApplication,

	// repository
	NewLinkStore,
	NewCQSLinkStore,
	NewQueryLinkStore,

	NewLinkService,
)

func InitLinkMQ(ctx context.Context, log logger.Logger, mq mq.MQ, service *link.Service) (*api_mq.Event, error) {
	linkMQ, err := api_mq.New(mq, log, service)
	if err != nil {
		return nil, err
	}

	return linkMQ, nil
}

func NewLinkStore(ctx context.Context, log logger.Logger, db db.DB, cache *cache.Cache) (crud.Repository, error) {
	linkStore, err := crud.New(ctx, log, db, cache)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

func NewCQSLinkStore(ctx context.Context, log logger.Logger, db db.DB, cache *cache.Cache) (*cqs.Store, error) {
	store, err := cqs.New(ctx, log, db, cache)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func NewQueryLinkStore(ctx context.Context, log logger.Logger, db db.DB, cache *cache.Cache) (*query.Store, error) {
	store, err := query.New(ctx, log, db, cache)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func NewLinkApplication(log logger.Logger, mq mq.MQ, metadataService metadata_rpc.MetadataServiceClient, store crud.Repository, authPermission *authzed.Client) (*link.Service, error) {
	linkService, err := link.New(log, mq, metadataService, store, authPermission)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewLinkCQRSApplication(log logger.Logger, cqsStore *cqs.Store, queryStore *query.Store) (*link_cqrs.Service, error) {
	linkCQRSService, err := link_cqrs.New(log, cqsStore, queryStore)
	if err != nil {
		return nil, err
	}

	return linkCQRSService, nil
}

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (link_rpc.LinkServiceClient, error) {
	LinkServiceClient := link_rpc.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewSitemapApplication(log logger.Logger, dataBus mq.MQ) (*sitemap.Service, error) {
	sitemapService, err := sitemap.New(log, dataBus)
	if err != nil {
		return nil, err
	}

	return sitemapService, nil
}

func NewLinkCQRSRPCServer(runRPCServer *rpc.Server, application *link_cqrs.Service, log logger.Logger) (*cqrs.Link, error) {
	linkRPCServer, err := cqrs.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewLinkRPCServer(runRPCServer *rpc.Server, application *link.Service, log logger.Logger) (*link_rpc.Link, error) {
	linkRPCServer, err := link_rpc.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewSitemapRPCServer(runRPCServer *rpc.Server, application *sitemap.Service, log logger.Logger) (*sitemap_rpc.Sitemap, error) {
	sitemapRPCServer, err := sitemap_rpc.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return sitemapRPCServer, nil
}

func NewRunRPCServer(runRPCServer *rpc.Server, _ *cqrs.Link, _ *link_rpc.Link) (*run.Response, error) {
	return run.Run(runRPCServer)
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (metadata_rpc.MetadataServiceClient, error) {
	metadataRPCClient := metadata_rpc.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewLinkService(
	// Common
	log logger.Logger,
	config *config.Config,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Security
	authPermission *authzed.Client,

	// Application
	linkService *link.Service,
	linkCQRSService *link_cqrs.Service,
	sitemapService *sitemap.Service,

	// Delivery
	linkMQ *api_mq.Event,
	run *run.Response,
	linkRPCServer *link_rpc.Link,
	linkCQRSRPCServer *cqrs.Link,
	sitemapRPCServer *sitemap_rpc.Sitemap,

	// Repository
	linkStore crud.Repository,

	// CQRS Repository
	cqsStore *cqs.Store,
	queryStore *query.Store,
) (*LinkService, error) {
	return &LinkService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		// Security
		authPermission: authPermission,

		// Application
		linkService:     linkService,
		linkCQRSService: linkCQRSService,
		sitemapService:  sitemapService,

		// Delivery
		run:               run,
		linkRPCServer:     linkRPCServer,
		linkCQRSRPCServer: linkCQRSRPCServer,
		sitemapRPCServer:  sitemapRPCServer,
		linkMQ:            linkMQ,

		// Repository
		linkStore: linkStore,

		// CQRS
		cqsStore:   cqsStore,
		queryStore: queryStore,
	}, nil
}

func InitializeLinkService() (*LinkService, func(), error) {
	panic(wire.Build(LinkSet))
}
