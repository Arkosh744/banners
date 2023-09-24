package models

type CreateRequest struct {
	Description string
}

type BannerSlotRequest struct {
	SlotID   int64 `json:"slot_id"`
	BannerID int64 `json:"banner_id"`
}

type EventRequest struct {
	SlotID   int64 `json:"slot_id"`
	BannerID int64 `json:"banner_id"`
	GroupID  int64 `json:"group_id"`
}

type NextBannerRequest struct {
	SlotID  int64 `json:"slot_id"`
	GroupID int64 `json:"group_id"`
}
