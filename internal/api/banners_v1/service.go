package banners_v1

import (
	"context"
	desc "github.com/Arkosh744/banners/pkg/banners_v1"
	"github.com/Arkosh744/banners/pkg/models"
)

type Implementation struct {
	desc.UnimplementedBannersServer

	bannersService Service
}

func NewImplementation(s Service) *Implementation {
	return &Implementation{
		bannersService: s,
	}
}

type Service interface {
	CreateSlot(context.Context, *models.CreateRequest) (*models.Slot, error)
	CreateBanner(context.Context, *models.CreateRequest) (*models.Banner, error)
	CreateGroup(context.Context, *models.CreateRequest) (*models.Group, error)
	AddBannerToSlot(context.Context, *models.BannerSlotRequest) error
	DeleteBannerFromSlot(context.Context, *models.BannerSlotRequest) error
	CreateClickEvent(context.Context, *models.EventRequest) error
	NextBanner(context.Context, *models.NextBannerRequest) (int64, error)
}
