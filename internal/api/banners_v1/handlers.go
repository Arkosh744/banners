package banners_v1

import (
	"context"
	desc "github.com/Arkosh744/banners/pkg/banners_v1"
	"github.com/Arkosh744/banners/pkg/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) CreateSlot(ctx context.Context, req *desc.CreateSlotReq) (*desc.SlotResp, error) {
	res, err := i.bannersService.CreateSlot(ctx, &models.CreateRequest{
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error create slot: %v", err)
	}

	return &desc.SlotResp{
		Id:          res.ID,
		Description: res.Description,
	}, nil
}

func (i *Implementation) CreateBanner(ctx context.Context, req *desc.CreateBannerReq) (*desc.BannerResp, error) {
	res, err := i.bannersService.CreateBanner(ctx, &models.CreateRequest{
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error create banner: %v", err)
	}

	return &desc.BannerResp{
		Id:          res.ID,
		Description: res.Description,
	}, nil
}

func (i *Implementation) CreateGroup(ctx context.Context, req *desc.CreateGroupReq) (*desc.GroupResp, error) {
	res, err := i.bannersService.CreateGroup(ctx, &models.CreateRequest{
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error create group: %v", err)
	}

	return &desc.GroupResp{
		Id:          res.ID,
		Description: res.Description,
	}, nil
}

func (i *Implementation) AddBannerToSlot(ctx context.Context, req *desc.BannerSlotRequest) (*emptypb.Empty, error) {
	if err := i.bannersService.AddBannerToSlot(ctx, &models.BannerSlotRequest{
		SlotID:   req.GetSlotId(),
		BannerID: req.GetBannerId(),
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error add banner to slot: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) DeleteBannerFromSlot(ctx context.Context, req *desc.BannerSlotRequest) (*emptypb.Empty, error) {
	if err := i.bannersService.DeleteBannerFromSlot(ctx, &models.BannerSlotRequest{
		SlotID:   req.GetSlotId(),
		BannerID: req.GetBannerId(),
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error delete banner from slot: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) CreateClickEvent(ctx context.Context, req *desc.ClickEvent) (*emptypb.Empty, error) {
	if err := i.bannersService.CreateClickEvent(ctx, &models.EventRequest{
		SlotID:   req.GetSlotId(),
		BannerID: req.GetBannerId(),
		GroupID:  req.GetGroupId(),
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "error create click event: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) NextBanner(ctx context.Context, req *desc.NextBannerRequest) (*desc.NextBannerResponse, error) {
	res, err := i.bannersService.NextBanner(ctx, &models.NextBannerRequest{
		SlotID:  req.GetSlotId(),
		GroupID: req.GetGroupId(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error next banner: %v", err)
	}

	return &desc.NextBannerResponse{BannerId: res}, nil
}
