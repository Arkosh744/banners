package app

import (
	"context"
	bannersV1 "github.com/Arkosh744/banners/internal/api/banners_v1"
	"github.com/Arkosh744/banners/internal/config"
	"github.com/Arkosh744/banners/internal/log"
	"github.com/Arkosh744/banners/internal/service"
	"github.com/Arkosh744/banners/pkg/closer"
	"github.com/Arkosh744/banners/pkg/pg"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type serviceProvider struct {
	repo service.Repository

	pgClient pg.Client
}

func newServiceProvider(_ context.Context) *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(config.AppConfig.GetPostgresDSN())
		if err != nil {
			log.Fatal("failed to parse pg config", zap.Error(err))
		}

		cl, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			log.Fatal("failed to get pg client", zap.Error(err))
		}

		if cl.PG().Ping(ctx) != nil {
			log.Fatal("failed to ping pg", zap.Error(err))
		}

		closer.Add(cl.Close)

		s.pgClient = cl
	}

	return s.pgClient
}

func (s *serviceProvider) GetRepo(ctx context.Context) service.Repository {

	return nil
}

func (s *serviceProvider) GetService(ctx context.Context) bannersV1.Service {

	return nil
}

func (s *serviceProvider) GetBannersImpl(ctx context.Context) *bannersV1.Implementation {

	return nil
}
