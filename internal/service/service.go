//go:generate mockgen -package=service -destination=./service_mock_internal_test.go -source=${GOFILE}
package service

import (
	"context"
	"fmt"

	"github.com/Arkosh744/banners/pkg/algo"
	"github.com/Arkosh744/banners/pkg/models"
)

type Service struct {
	repo   Repository
	broker MessageBroker
}

func New(repo Repository, broker MessageBroker) *Service {
	return &Service{
		repo:   repo,
		broker: broker,
	}
}

type Repository interface {
	CreateSlot(ctx context.Context, description string) (int, error)
	CreateBanner(ctx context.Context, description string) (int, error)
	CreateSocGroup(ctx context.Context, description string) (int, error)
	AddBannerToSlot(ctx context.Context, req *models.BannerSlotRequest) error
	DeleteBannerSlot(ctx context.Context, req *models.BannerSlotRequest) error
	CreateClickEvent(ctx context.Context, req *models.EventRequest) error
	GetBannersInfo(ctx context.Context, req *models.NextBannerRequest) ([]models.BannerStats, error)
	IncrementBannerView(ctx context.Context, req *models.EventRequest) error
}

type MessageBroker interface {
	SendMessage(BannerID, SlotID, GroupID int64, msgType string) error
}

func (s *Service) CreateSlot(ctx context.Context, req *models.CreateRequest) (*models.Slot, error) {
	res, err := s.repo.CreateSlot(ctx, req.Description)
	if err != nil {
		return nil, fmt.Errorf("error create slot: %w", err)
	}

	return &models.Slot{
		ID:          int64(res),
		Description: req.Description,
	}, nil
}

func (s *Service) CreateBanner(ctx context.Context, req *models.CreateRequest) (*models.Banner, error) {
	res, err := s.repo.CreateBanner(ctx, req.Description)
	if err != nil {
		return nil, fmt.Errorf("error create banner: %w", err)
	}

	return &models.Banner{
		ID:          int64(res),
		Description: req.Description,
	}, nil
}

func (s *Service) CreateGroup(ctx context.Context, req *models.CreateRequest) (*models.Group, error) {
	res, err := s.repo.CreateSocGroup(ctx, req.Description)
	if err != nil {
		return nil, fmt.Errorf("create soc group: %w", err)
	}

	return &models.Group{
		ID:          int64(res),
		Description: req.Description,
	}, nil
}

func (s *Service) AddBannerToSlot(ctx context.Context, req *models.BannerSlotRequest) error {
	if err := s.repo.AddBannerToSlot(ctx, req); err != nil {
		return fmt.Errorf("error add banner to slot: %w", err)
	}

	return nil
}

func (s *Service) DeleteBannerFromSlot(ctx context.Context, req *models.BannerSlotRequest) error {
	if err := s.repo.DeleteBannerSlot(ctx, req); err != nil {
		return fmt.Errorf("error delete banner from slot: %w", err)
	}

	return nil
}

func (s *Service) CreateClickEvent(ctx context.Context, req *models.EventRequest) error {
	if err := s.repo.CreateClickEvent(ctx, req); err != nil {
		return fmt.Errorf("error create click event: %w", err)
	}

	if err := s.broker.SendMessage(req.BannerID, req.SlotID, req.GroupID, models.KafkaTypeClick); err != nil {
		return fmt.Errorf("error send notification message click: %w", err)
	}

	return nil
}

func (s *Service) NextBanner(ctx context.Context, req *models.NextBannerRequest) (int64, error) {
	banners, err := s.repo.GetBannersInfo(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("error get banners info: %w", err)
	}

	bannerID, err := algo.MultiArmedBandit(banners)
	if err != nil {
		return 0, fmt.Errorf("error get banner id: %w", err)
	}

	if err = s.broker.SendMessage(bannerID, req.SlotID, req.GroupID, models.KafkaTypeView); err != nil {
		return 0, fmt.Errorf("error send notification message view: %w", err)
	}

	if err = s.repo.IncrementBannerView(ctx, &models.EventRequest{
		SlotID:   req.SlotID,
		BannerID: bannerID,
		GroupID:  req.GroupID,
	}); err != nil {
		return 0, fmt.Errorf("error increment banner view: %w", err)
	}

	return bannerID, nil
}
