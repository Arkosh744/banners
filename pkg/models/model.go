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
	BannerID   int64  `json:"bannerId"`
	SlotID     int64  `json:"slotId"`
	GroupID    *int64 `json:"groupId"`
	ViewCount  int64  `json:"viewCount"`
	ClickCount int64  `json:"clickCount"`
}
