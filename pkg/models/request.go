package models

type CreateRequest struct {
	Description string
}

type BannerSlotRequest struct {
	SlotID   int64 `json:"slotId"`
	BannerID int64 `json:"bannerId"`
}

type EventRequest struct {
	SlotID   int64 `json:"slotId"`
	BannerID int64 `json:"bannerId"`
	GroupID  int64 `json:"groupId"`
}

type NextBannerRequest struct {
	SlotID  int64 `json:"slotId"`
	GroupID int64 `json:"groupId"`
}
