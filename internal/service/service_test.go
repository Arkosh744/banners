package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/Arkosh744/banners/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestService_CreateBanner(t *testing.T) { //nolint: dupl // test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockKafka := NewMockKafka(ctrl)

	s := New(mockRepo, mockKafka)

	ctx := context.Background()

	errCreateBanner := errors.New("error create banner")
	tests := []struct {
		name      string
		req       *models.CreateRequest
		repoMock  func(m *MockRepository)
		kafkaMock func(m *MockKafka)
		want      *models.Banner
		wantErr   error
	}{
		{
			name: "Success",
			req:  &models.CreateRequest{Description: "test"},
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateBanner(ctx, "test").Return(1, nil).Times(1)
			},
			want: &models.Banner{ID: 1, Description: "test"},
		},
		{
			name:    "Error",
			req:     &models.CreateRequest{Description: "test"},
			wantErr: errCreateBanner,
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateBanner(ctx, "test").Return(0, errCreateBanner).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}

			got, err := s.CreateBanner(ctx, tt.req)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBanner() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateGroup(t *testing.T) { //nolint: dupl // test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockKafka := NewMockKafka(ctrl)

	s := New(mockRepo, mockKafka)

	ctx := context.Background()

	errCreateBanner := errors.New("error create group")
	tests := []struct {
		name      string
		req       *models.CreateRequest
		repoMock  func(m *MockRepository)
		kafkaMock func(m *MockKafka)
		want      *models.Group
		wantErr   error
	}{
		{
			name: "Success",
			req:  &models.CreateRequest{Description: "test"},
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateSocGroup(ctx, "test").Return(1, nil).Times(1)
			},
			want: &models.Group{ID: 1, Description: "test"},
		},
		{
			name:    "Error",
			req:     &models.CreateRequest{Description: "test"},
			wantErr: errCreateBanner,
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateSocGroup(ctx, "test").Return(0, errCreateBanner).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}

			got, err := s.CreateGroup(ctx, tt.req)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSocGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateSlot(t *testing.T) { //nolint: dupl // test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockKafka := NewMockKafka(ctrl)

	s := New(mockRepo, mockKafka)

	ctx := context.Background()

	errCreateBanner := errors.New("error create slot")
	tests := []struct {
		name      string
		req       *models.CreateRequest
		repoMock  func(m *MockRepository)
		kafkaMock func(m *MockKafka)
		want      *models.Slot
		wantErr   error
	}{
		{
			name: "Success",
			req:  &models.CreateRequest{Description: "test"},
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateSlot(ctx, "test").Return(1, nil).Times(1)
			},
			want: &models.Slot{ID: 1, Description: "test"},
		},
		{
			name:    "Error",
			req:     &models.CreateRequest{Description: "test"},
			wantErr: errCreateBanner,
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateSlot(ctx, "test").Return(0, errCreateBanner).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}

			got, err := s.CreateSlot(ctx, tt.req)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSlot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddBannerToSlot(t *testing.T) { //nolint: dupl // test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockKafka := NewMockKafka(ctrl)

	s := New(mockRepo, mockKafka)

	ctx := context.Background()

	request := &models.BannerSlotRequest{
		SlotID:   1,
		BannerID: 1,
	}
	testErr := errors.New("error")
	tests := []struct {
		name      string
		req       *models.BannerSlotRequest
		repoMock  func(m *MockRepository)
		kafkaMock func(m *MockKafka)
		wantErr   error
	}{
		{
			name: "Success",
			req:  request,
			repoMock: func(m *MockRepository) {
				m.EXPECT().AddBannerToSlot(ctx, request).Return(nil).Times(1)
			},
		},
		{
			name:    "Error",
			req:     request,
			wantErr: testErr,
			repoMock: func(m *MockRepository) {
				m.EXPECT().AddBannerToSlot(ctx, request).Return(testErr).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}

			err := s.AddBannerToSlot(ctx, request)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
			}
		})
	}
}

func TestService_DeleteBannerFromSlot(t *testing.T) { //nolint: dupl // test
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockKafka := NewMockKafka(ctrl)

	s := New(mockRepo, mockKafka)

	ctx := context.Background()

	request := &models.BannerSlotRequest{
		SlotID:   1,
		BannerID: 1,
	}
	testErr := errors.New("error")
	tests := []struct {
		name      string
		req       *models.BannerSlotRequest
		repoMock  func(m *MockRepository)
		kafkaMock func(m *MockKafka)
		wantErr   error
	}{
		{
			name: "Success",
			req:  request,
			repoMock: func(m *MockRepository) {
				m.EXPECT().DeleteBannerSlot(ctx, request).Return(nil).Times(1)
			},
		},
		{
			name:    "Error",
			req:     request,
			wantErr: testErr,
			repoMock: func(m *MockRepository) {
				m.EXPECT().DeleteBannerSlot(ctx, request).Return(testErr).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}

			err := s.DeleteBannerFromSlot(ctx, request)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
			}
		})
	}
}

func TestService_CreateClickEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockKafka := NewMockKafka(ctrl)

	s := New(mockRepo, mockKafka)

	ctx := context.Background()

	request := &models.EventRequest{
		BannerID: 1,
		SlotID:   1,
		GroupID:  1,
	}
	testErr := errors.New("error")
	tests := []struct {
		name      string
		req       *models.EventRequest
		repoMock  func(m *MockRepository)
		kafkaMock func(m *MockKafka)
		wantErr   error
	}{
		{
			name: "Success",
			req:  request,
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateClickEvent(ctx, request).Return(nil).Times(1)
			},
			kafkaMock: func(m *MockKafka) {
				m.EXPECT().SendMessage(int64(1), int64(1), int64(1), models.KafkaTypeClick).Return(nil)
			},
		},
		{
			name:    "Error",
			req:     request,
			wantErr: testErr,
			repoMock: func(m *MockRepository) {
				m.EXPECT().CreateClickEvent(ctx, request).Return(testErr).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}
			if tt.kafkaMock != nil {
				tt.kafkaMock(mockKafka)
			}

			err := s.CreateClickEvent(ctx, request)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
			}
		})
	}
}

func TestService_NextBanner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepository(ctrl)
	mockKafka := NewMockKafka(ctrl)

	s := New(mockRepo, mockKafka)

	ctx := context.Background()

	request := &models.NextBannerRequest{
		SlotID:  1,
		GroupID: 1,
	}
	testErr := errors.New("error")
	tests := []struct {
		name      string
		req       *models.NextBannerRequest
		repoMock  func(m *MockRepository)
		kafkaMock func(m *MockKafka)
		want      int64
		wantErr   error
	}{
		{
			name: "Success",
			req:  request,
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetBannersInfo(ctx, request).Return([]models.BannerStats{
					{
						BannerID:   1,
						SlotID:     1,
						GroupID:    lo.ToPtr(int64(1)),
						ViewCount:  1,
						ClickCount: 1,
					},
				}, nil).Times(1)
				m.EXPECT().IncrementBannerView(ctx, &models.EventRequest{
					SlotID:   int64(1),
					BannerID: int64(1),
					GroupID:  int64(1),
				}).Return(nil).Times(1)
			},
			kafkaMock: func(m *MockKafka) {
				m.EXPECT().SendMessage(int64(1), int64(1), int64(1), models.KafkaTypeView).Return(nil)
			},
			want: int64(1),
		},
		{
			name:    "Error",
			req:     request,
			wantErr: testErr,
			repoMock: func(m *MockRepository) {
				m.EXPECT().GetBannersInfo(ctx, request).Return(nil, testErr).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.repoMock != nil {
				tt.repoMock(mockRepo)
			}

			if tt.kafkaMock != nil {
				tt.kafkaMock(mockKafka)
			}

			bannerID, err := s.NextBanner(ctx, request)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tt.wantErr))
			}

			if err == nil && tt.want != 0 {
				require.Equal(t, tt.want, bannerID)
			}
		})
	}
}
