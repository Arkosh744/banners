package service

import (
	"context"
	"github.com/Arkosh744/banners/internal/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	repo  Repository
	kafka Kafka
}

func New(repo Repository, kafka Kafka) *Service {
	return &Service{
		repo:  repo,
		kafka: kafka,
	}
}

type Repository interface {
	CreateSlot(ctx context.Context, description string) (int, error)
	CreateBanner(ctx context.Context, description string) (int, error)
	CreateSocGroup(ctx context.Context, description string) (int, error)
}

type Kafka interface {
}

func (s *Service) CreateSlot(ctx context.Context, req *models.CreateRequest) (*models.Slot, error) {

	return nil, nil
}

func (s *Service) CreateBanner(ctx context.Context, req *models.CreateRequest) (*models.Banner, error) {

	return nil, nil
}

func (s *Service) CreateGroup(ctx context.Context, req *models.CreateRequest) (*models.Group, error) {

	return nil, nil
}

func (s *Service) AddBannerToSlot(ctx context.Context, req *models.BannerSlotRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (s *Service) DeleteBannerFromSlot(ctx context.Context, req *models.BannerSlotRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (s *Service) CreateClickEvent(ctx context.Context, req *models.ClickEventRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (s *Service) NextBanner(ctx context.Context, req *models.NextBannerRequest) (*models.Banner, error) {

	return nil, nil
}
