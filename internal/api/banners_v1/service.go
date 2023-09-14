package banners_v1

import (
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
	CreateSlot(context.Context, *CreateSlotReq) (*SlotResp, error)
	CreateBanner(context.Context, *CreateBannerReq) (*BannerResp, error)
	CreateGroup(context.Context, *CreateGroupReq) (*GroupResp, error)
	AddBannerToSlot(context.Context, *BannerSlotRequest) (*emptypb.Empty, error)
	DeleteBannerFromSlot(context.Context, *BannerSlotRequest) (*emptypb.Empty, error)
	CreateClickEvent(context.Context, *ClickEvent) (*emptypb.Empty, error)
	NextBanner(context.Context, *NextBannerRequest) (*BannerResp, error)
}
