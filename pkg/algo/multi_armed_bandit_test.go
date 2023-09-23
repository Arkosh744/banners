package algo_test

import (
	"testing"

	"github.com/Arkosh744/banners/pkg/algo"
	"github.com/Arkosh744/banners/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestMultiArmedBandit(t *testing.T) {
	var testCases = []struct {
		name             string
		banners          []models.BannerStats
		expectedBannerID []int
		expectedError    error
	}{
		{
			name:             "no banners",
			banners:          []models.BannerStats{},
			expectedBannerID: []int{0},
			expectedError:    models.ErrGetBanner,
		},
		{
			name:             "single banner",
			banners:          []models.BannerStats{{BannerID: 1, ViewCount: 50, ClickCount: 10}},
			expectedBannerID: []int{1},
			expectedError:    nil,
		},
		{
			name:             "multiple banners same performance",
			banners:          []models.BannerStats{{BannerID: 1, ViewCount: 50, ClickCount: 10}, {BannerID: 2, ViewCount: 50, ClickCount: 10}},
			expectedBannerID: []int{1, 2},
			expectedError:    nil,
		},
		{
			name:             "multiple banners different performance",
			banners:          []models.BannerStats{{BannerID: 1, ViewCount: 50, ClickCount: 10}, {BannerID: 2, ViewCount: 50, ClickCount: 5}},
			expectedBannerID: []int{1},
			expectedError:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bannerID, err := algo.MultiArmedBandit(tt.banners)
			require.Equal(t, tt.expectedError, err)
			require.Contains(t, tt.expectedBannerID, bannerID)
		})
	}
}
