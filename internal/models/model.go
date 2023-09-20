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
