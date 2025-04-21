package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/usecase"
)

var (
	ErrNotPublish = errors.New("event not published")
	ErrNotFound   = errors.New("user not found")
	ErrCreate     = errors.New("user not created")
	ErrUpdate     = errors.New("user not updated")
	ErrDelete     = errors.New("user not deleted")
	ErrCacheGet   = errors.New("user not found in cache")
	ErrCacheSet   = errors.New("user not saved in cache")
	ErrCacheDel   = errors.New("user not deleted from cache")

	//nolint:gochecknoglobals
	user = &entity.User{
		Name:  "Test",
		Email: "1@1.com",
		Role:  "admin",
	}
)

func NewUseCase(t *testing.T) (*usecase.UseCase, *MockMessageSys, *MockUserRepo, *MockUserCache) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ms := NewMockMessageSys(ctrl)
	ur := NewMockUserRepo(ctrl)
	uc := NewMockUserCache(ctrl)

	return usecase.New(ur, uc, ms), ms, ur, uc
}

func TestGetUsersNotPublish(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUsers")).Return(ErrNotPublish)

	_, err := ucase.GetUsers(t.Context())
	require.Error(t, err)
}

func TestGetUsersNotFound(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUsers")).Return(nil)
	repo.EXPECT().GetUsers(t.Context()).Return(nil, ErrNotFound)

	_, err := ucase.GetUsers(t.Context())
	require.Error(t, err)
}

func TestGetUsers(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUsers")).Return(nil)
	repo.EXPECT().GetUsers(t.Context()).Return([]*entity.User{user}, nil)

	result, _ := ucase.GetUsers(t.Context())
	require.Equal(t, []*entity.User{user}, result)
}

func TestGetUserByIDNotPublish(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUserByID")).Return(ErrNotPublish)

	_, err := ucase.GetUserByID(t.Context(), "ID")
	require.Error(t, err)
}

func TestGetUserByIDCacheGet(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUserByID")).Return(nil)
	cache.EXPECT().Get(t.Context(), "ID").Return(nil, ErrCacheGet)

	_, err := ucase.GetUserByID(t.Context(), "ID")
	require.Error(t, err)
}

func TestGetUserByIDNotFound(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUserByID")).Return(nil)
	cache.EXPECT().Get(t.Context(), "ID").Return(nil, nil)
	repo.EXPECT().GetUserByID(t.Context(), "ID").Return(nil, ErrNotFound)

	result, err := ucase.GetUserByID(t.Context(), "ID")
	require.Nil(t, result)
	require.Error(t, err)
}

func TestGetUserByIDCacheSet(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUserByID")).Return(nil)
	cache.EXPECT().Get(t.Context(), "ID").Return(nil, nil)
	repo.EXPECT().GetUserByID(t.Context(), "ID").Return(user, nil)
	cache.EXPECT().Set(t.Context(), user, time.Minute).Return(ErrCacheSet)

	_, err := ucase.GetUserByID(t.Context(), "ID")
	require.Error(t, err)
}

func TestGetUserByIDCacheFound(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUserByID")).Return(nil)
	cache.EXPECT().Get(t.Context(), "ID").Return(user, nil)

	result, err := ucase.GetUserByID(t.Context(), "ID")
	require.Equal(t, user, result)
	require.NoError(t, err)
}

func TestGetUserByIDRepoFound(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.GetUserByID")).Return(nil)
	cache.EXPECT().Get(t.Context(), "ID").Return(nil, nil)
	repo.EXPECT().GetUserByID(t.Context(), "ID").Return(user, nil)
	cache.EXPECT().Set(t.Context(), user, time.Minute).Return(nil)

	result, err := ucase.GetUserByID(t.Context(), "ID")
	require.Equal(t, user, result)
	require.NoError(t, err)
}

func TestCreateNotPublish(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Create")).Return(ErrNotPublish)

	err := ucase.Create(t.Context(), user)
	require.Error(t, err)
}

func TestCreateError(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Create")).Return(nil)
	repo.EXPECT().Create(t.Context(), user).Return(ErrCreate)

	err := ucase.Create(t.Context(), user)
	require.Error(t, err)
}

func TestCreateCacheSet(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Create")).Return(nil)
	repo.EXPECT().Create(t.Context(), user).Return(nil)
	cache.EXPECT().Set(t.Context(), user, time.Minute).Return(ErrCacheSet)

	err := ucase.Create(t.Context(), user)
	require.Error(t, err)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Create")).Return(nil)
	repo.EXPECT().Create(t.Context(), user).Return(nil)
	cache.EXPECT().Set(t.Context(), user, time.Minute).Return(nil)

	err := ucase.Create(t.Context(), user)
	require.NoError(t, err)
}

func TestUpdateNotPublish(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Update")).Return(ErrNotPublish)

	err := ucase.Update(t.Context(), user)
	require.Error(t, err)
}

func TestUpdateError(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Update")).Return(nil)
	repo.EXPECT().Update(t.Context(), user).Return(ErrUpdate)

	err := ucase.Update(t.Context(), user)
	require.Error(t, err)
}

func TestUpdateCacheSet(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Update")).Return(nil)
	repo.EXPECT().Update(t.Context(), user).Return(nil)
	cache.EXPECT().Set(t.Context(), user, time.Minute).Return(ErrCacheSet)

	err := ucase.Update(t.Context(), user)
	require.Error(t, err)
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)

	mbus.EXPECT().Publish("methods", []byte("usecase.Update")).Return(nil)
	repo.EXPECT().Update(t.Context(), user).Return(nil)
	cache.EXPECT().Set(t.Context(), user, time.Minute).Return(nil)

	err := ucase.Update(t.Context(), user)
	require.NoError(t, err)
}

func TestDeleteNotPublish(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, _ := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Delete")).Return(ErrNotPublish)

	err := ucase.Delete(t.Context(), "ID")
	require.Error(t, err)
}

func TestDeleteCacheDel(t *testing.T) {
	t.Parallel()

	ucase, mbus, _, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Delete")).Return(nil)
	cache.EXPECT().Del(t.Context(), "ID").Return(ErrCacheDel)

	err := ucase.Delete(t.Context(), "ID")
	require.Error(t, err)
}

func TestDeleteError(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Delete")).Return(nil)
	cache.EXPECT().Del(t.Context(), "ID").Return(nil)
	repo.EXPECT().Delete(t.Context(), "ID").Return(ErrDelete)

	err := ucase.Delete(t.Context(), "ID")
	require.Error(t, err)
}

func TestDelete(t *testing.T) {
	t.Parallel()

	ucase, mbus, repo, cache := NewUseCase(t)
	mbus.EXPECT().Publish("methods", []byte("usecase.Delete")).Return(nil)
	cache.EXPECT().Del(t.Context(), "ID").Return(nil)
	repo.EXPECT().Delete(t.Context(), "ID").Return(nil)

	err := ucase.Delete(t.Context(), "ID")
	require.NoError(t, err)
}
