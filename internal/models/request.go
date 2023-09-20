package models

type CreateRequest struct {
	Description string
}

type BannerSlotRequest struct {
	SlotID   int64 `json:"slotID"`
	BannerID int64 `json:"bannerID"`
}

type ClickEventRequest struct {
	SlotID   int64 `json:"slotID"`
	BannerID int64 `json:"bannerID"`
	GroupID  int64 `json:"groupID"`
}

type NextBannerRequest struct {
	SlotID  int64 `json:"slotID"`
	GroupID int64 `json:"groupID"`
}
