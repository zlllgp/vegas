package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zlllgp/vegas/internal/domain/repository/mock_repository"
	"github.com/zlllgp/vegas/pkg/errno"
	"os"
	"testing"
)

var (
	mockCtrl      *gomock.Controller
	mockRmbRepo   *mock_repository.MockRmbRepository
	mockActRepo   *mock_repository.MockActivityRepository
	mockRedisRepo *mock_repository.MockRedisActivityRepository
	mockCacheRepo *mock_repository.MockCacheRepository
	drawService   *DrawServiceImpl
)

func TestMain(m *testing.M) {
	mockCtrl = gomock.NewController(nil)
	mockRmbRepo = mock_repository.NewMockRmbRepository(mockCtrl)
	mockActRepo = mock_repository.NewMockActivityRepository(mockCtrl)
	mockRedisRepo = mock_repository.NewMockRedisActivityRepository(mockCtrl)
	mockCacheRepo = mock_repository.NewMockCacheRepository(mockCtrl)
	drawService = NewDrawServiceImpl(mockRmbRepo, mockActRepo, mockRedisRepo, mockCacheRepo)
	exitCode := m.Run()

	mockCtrl.Finish()
	os.Exit(exitCode)
}

func TestDrawService_Draw(t *testing.T) {
	mockRmbRepo.EXPECT().IsSafe(context.Background(), int64(1)).Return(false)

	result, err := drawService.Draw(context.Background(), int64(1), "1")

	assert.Nil(t, result)
	assert.ErrorIs(t, err, errno.DrawErr)
}
