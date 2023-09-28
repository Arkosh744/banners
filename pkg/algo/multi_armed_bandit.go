package algo

import (
	"math"
	"math/rand"

	"github.com/Arkosh744/banners/pkg/models"
)

func MultiArmedBandit(banners []models.BannerStats) (int64, error) {
	if len(banners) == 0 {
		return 0, models.ErrNoBanner
	}

	if len(banners) == 1 {
		return banners[0].BannerID, nil
	}

	var (
		bannerID   int64
		totalViews int64
		maxIncome  float64 = -1
	)

	for _, banner := range banners {
		totalViews += banner.ViewCount
	}

	bannerIds := make([]int64, 0, len(banners))
	for _, banner := range banners {
		bannerIncome := (float64(banner.ClickCount) / float64(banner.ViewCount)) +
			math.Sqrt((2.0*math.Log(float64(totalViews)))/float64(banner.ViewCount))

		if bannerIncome < maxIncome {
			continue
		}

		if bannerIncome > maxIncome {
			maxIncome = bannerIncome
			bannerIds = []int64{}
		}

		bannerIds = append(bannerIds, banner.BannerID)
	}

	if len(bannerIds) == 0 {
		return 0, models.ErrGetBanner
	}

	if len(bannerIds) == 1 {
		return bannerIds[0], nil
	}

	bannerID = bannerIds[rand.Intn(len(bannerIds))] //nolint: gosec // it's ok here

	return bannerID, nil
}
