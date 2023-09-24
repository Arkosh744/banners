package models

const (
	KafkaTypeView  = "view"
	KafkaTypeClick = "click"
)

type Slot struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
}

type Banner struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
}

type Group struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
}

type BannerStats struct {
	BannerID   int64  `json:"banner_id"`
	SlotID     int64  `json:"slot_id"`
	GroupID    *int64 `json:"group_id"`
	ViewCount  int64  `json:"view_count"`
	ClickCount int64  `json:"click_count"`
}
