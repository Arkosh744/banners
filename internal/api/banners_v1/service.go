package banners_v1

import (
	"context"
	"github.com/Arkosh744/banners/internal/models"
	desc "github.com/Arkosh744/banners/pkg/banners_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	desc.UnimplementedBannersServer

	cartService Service
}

func NewImplementation(s Service) *Implementation {
	return &Implementation{
		cartService: s,
	}
}

type Service interface {
	CreateSlot(context.Context, *models.CreateRequest) (*models.Slot, error)
	CreateBanner(context.Context, *models.CreateRequest) (*models.Banner, error)
	CreateGroup(context.Context, *models.CreateRequest) (*models.Group, error)
	AddBannerToSlot(context.Context, *models.BannerSlotRequest) (*emptypb.Empty, error)
	DeleteBannerFromSlot(context.Context, *models.BannerSlotRequest) (*emptypb.Empty, error)
	CreateClickEvent(context.Context, *models.ClickEventRequest) (*emptypb.Empty, error)
	NextBanner(context.Context, *models.NextBannerRequest) (*models.Banner, error)
}
