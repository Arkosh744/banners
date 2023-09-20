package app

import (
	"context"
	bannersV1 "github.com/Arkosh744/banners/internal/api/banners_v1"
	"github.com/Arkosh744/banners/internal/config"
	"github.com/Arkosh744/banners/internal/log"
	"github.com/Arkosh744/banners/internal/repo"
	"github.com/Arkosh744/banners/internal/service"
	"github.com/Arkosh744/banners/pkg/closer"
	"github.com/Arkosh744/banners/pkg/kafka"
	"github.com/Arkosh744/banners/pkg/pg"

	"github.com/IBM/sarama"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type serviceProvider struct {
	service     bannersV1.Service
	bannersImlp *bannersV1.Implementation

	pgClient pg.Client
	repo     service.Repository
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
	if s.repo == nil {
		s.repo = repo.NewRepo(s.GetPGClient(ctx))
	}

	return s.repo
}

func (s *serviceProvider) GetKafkaSyncProducer() (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()

	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Idempotent = true
	cfg.Net.MaxOpenRequests = 1

	producer, err := sarama.NewSyncProducer(config.AppConfig.Kafka.Brokers, cfg)
	if err != nil {
		return nil, err
	}

	log.Info("kafka sync producer created")

	admin, err := sarama.NewClusterAdmin(config.AppConfig.Kafka.Brokers, cfg)
	if err != nil {
		return nil, err
	}

	err = ensureTopicExists(admin, config.AppConfig.Kafka.Topic, 1, 1)
	if err != nil {
		return nil, err
	}

	closer.Add(producer.Close)

	return producer, nil
}

func ensureTopicExists(admin sarama.ClusterAdmin, topic string, numPartitions int32, replicationFactor int16) error {
	topics, err := admin.ListTopics()
	if err != nil {
		return err
	}

	if _, ok := topics[topic]; !ok {
		details := &sarama.TopicDetail{
			NumPartitions:     numPartitions,
			ReplicationFactor: replicationFactor,
		}

		if err = admin.CreateTopic(topic, details, false); err != nil {
			return err
		}
	}

	return nil
}

func (s *serviceProvider) GetService(ctx context.Context) bannersV1.Service {
	if s.service == nil {
		producer, err := s.GetKafkaSyncProducer()
		if err != nil {
			log.Fatal("failed to get kafka sync producer", zap.Error(err))
		}

		kafkaProducer := kafka.NewProducer(producer, config.AppConfig.Kafka.Topic)
		s.service = service.New(s.GetRepo(ctx), kafkaProducer)
	}

	return s.service
}

func (s *serviceProvider) GetBannersImpl(ctx context.Context) *bannersV1.Implementation {
	if s.bannersImlp == nil {
		s.bannersImlp = bannersV1.NewImplementation(s.GetService(ctx))
	}

	return nil
}
