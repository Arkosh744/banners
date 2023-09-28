package models

import "errors"

var (
	ErrGetBanner = errors.New("error get banner")
	ErrNoBanner  = errors.New("no banners found")
)
