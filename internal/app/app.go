package app

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/Arkosh744/banners/internal/config"
	"github.com/Arkosh744/banners/internal/log"
	descBannerV1 "github.com/Arkosh744/banners/pkg/banners_v1"
	"github.com/Arkosh744/banners/pkg/closer"
	"github.com/Arkosh744/banners/pkg/interceptor"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := app.RunGrpcServer(); err != nil {
			log.Fatal("failed to run grpc server", zap.Error(err))
		}
	}()

	wg.Wait()

	return nil
}

func (app *App) initDeps(ctx context.Context) error {
	for _, init := range []func(context.Context) error{
		config.Init,
		app.initLogger,
		app.initServiceProvider,
		app.initGrpcServer,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initLogger(ctx context.Context) error {
	if err := log.InitLogger(ctx, config.AppConfig.Log.Preset); err != nil {
		return err
	}

	return nil
}

func (app *App) initServiceProvider(ctx context.Context) error {
	app.serviceProvider = newServiceProvider(ctx)

	return nil
}

func (app *App) initGrpcServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			interceptor.LoggingInterceptor,
		),
		),
	)
	reflection.Register(app.grpcServer)

	descBannerV1.RegisterBannersServer(app.grpcServer, app.serviceProvider.GetBannersImpl(ctx))

	return nil
}

func (app *App) RunGrpcServer() error {
	log.Info(fmt.Sprintf("GRPC server listening on port %s", config.AppConfig.GetGRPCAddr()))

	list, err := net.Listen("tcp", config.AppConfig.GetGRPCAddr())
	if err != nil {
		return err
	}

	err = app.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
