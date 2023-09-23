package service

import (
	"context"
	"errors"
	"github.com/Arkosh744/banners/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestService_AddBannerToSlot(t *testing.T) {
	type fields struct {
		repo  Repository
		kafka Kafka
	}
	type args struct {
		ctx context.Context
		req *models.BannerSlotRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo:  tt.fields.repo,
				kafka: tt.fields.kafka,
			}
			if err := s.AddBannerToSlot(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("AddBannerToSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_CreateBanner(t *testing.T) {
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

func TestService_CreateGroup(t *testing.T) {
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

func TestService_CreateSlot(t *testing.T) {
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

func TestService_CreateClickEvent(t *testing.T) {
	type fields struct {
		repo  Repository
		kafka Kafka
	}
	type args struct {
		ctx context.Context
		req *models.EventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo:  tt.fields.repo,
				kafka: tt.fields.kafka,
			}
			if err := s.CreateClickEvent(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("CreateClickEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteBannerFromSlot(t *testing.T) {
	type fields struct {
		repo  Repository
		kafka Kafka
	}
	type args struct {
		ctx context.Context
		req *models.BannerSlotRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo:  tt.fields.repo,
				kafka: tt.fields.kafka,
			}
			if err := s.DeleteBannerFromSlot(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBannerFromSlot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_NextBanner(t *testing.T) {
	type fields struct {
		repo  Repository
		kafka Kafka
	}
	type args struct {
		ctx context.Context
		req *models.NextBannerRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo:  tt.fields.repo,
				kafka: tt.fields.kafka,
			}
			got, err := s.NextBanner(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NextBanner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NextBanner() got = %v, want %v", got, tt.want)
			}
		})
	}
}
