package banners_v1

import (
	"context"
	desc "github.com/Arkosh744/banners/pkg/banners_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CreateSlot(ctx context.Context, req *desc.CreateSlotReq) (*desc.SlotResp, error) {

	return nil, nil
}

func (i *Implementation) CreateBanner(ctx context.Context, req *desc.CreateBannerReq) (*desc.BannerResp, error) {

	return nil, nil
}

func (i *Implementation) CreateGroup(ctx context.Context, req *desc.CreateGroupReq) (*desc.GroupResp, error) {

	return nil, nil
}

func (i *Implementation) AddBannerToSlot(ctx context.Context, req *desc.BannerSlotRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (i *Implementation) DeleteBannerFromSlot(ctx context.Context, req *desc.BannerSlotRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (i *Implementation) CreateClickEvent(ctx context.Context, req *desc.ClickEvent) (*emptypb.Empty, error) {

	return nil, nil
}

func (i *Implementation) NextBanner(ctx context.Context, req *desc.NextBannerRequest) (*desc.BannerResp, error) {

	return nil, nil
}
