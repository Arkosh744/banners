package models

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
	BannerID          int64  `json:"bannerID"`
	BannerDescription string `json:"bannerDescription"`
	SlotID            int64  `json:"slotID"`
	SocialGroupID     int64  `json:"socialGroupID"`
	ViewCount         int64  `json:"viewCount"`
	ClickCount        int64  `json:"clickCount"`
}
