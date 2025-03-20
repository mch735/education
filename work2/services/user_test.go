package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mch735/education/work2/models/user"
	"github.com/mch735/education/work2/services"
	"github.com/mch735/education/work2/storages"
	"github.com/mch735/education/work2/storages/memory"
	"github.com/mch735/education/work2/storages/mock"
)

func TestUserServiceValidateError(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	_, err := service.CreateUser("Test", "1@1.com", "qwe")
	require.ErrorIs(t, err, services.ErrInvalidRole)

	_, err = service.CreateUser("Test", "example.com", "admin")
	require.ErrorIs(t, err, services.ErrInvalidEmail)

	_, err = service.CreateUser("", "1@1.com", "user")
	require.ErrorIs(t, err, services.ErrInvalidName)
}

func TestUserServiceSaveError(t *testing.T) {
	t.Parallel()

	repo := mock.NewMockErrorUserRepo()

	_, err := services.NewUserService(repo).CreateUser("Test", "1@1.com", "admin")
	require.ErrorIs(t, err, storages.ErrUserExist)
}

func TestUserServiceSave(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	record, _ := service.CreateUser("Test", "1@1.com", "user")
	require.Equal(t, "Test", record.Name)
	require.Equal(t, "1@1.com", record.Email)
	require.Equal(t, "user", record.Role)
}

func TestUserServiceGetUserError(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	_, err := service.GetUser("10")
	require.ErrorIs(t, err, storages.ErrUserNotFound)
}

func TestUserServiceGetUser(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	expect, err := service.CreateUser("Test", "1@1.com", "user")
	require.NoError(t, err)

	actual, err := service.GetUser(expect.ID)
	require.NoError(t, err)
	require.Equal(t, expect, actual)
}

func TestUserServiceRemoveUserError(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	err := service.RemoveUser("10")
	require.ErrorIs(t, err, storages.ErrUserNotFound)
}

func TestUserServiceRemoveUser(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	expect, err := service.CreateUser("Test", "1@1.com", "user")
	require.NoError(t, err)

	require.NoError(t, service.RemoveUser(expect.ID))

	_, err = service.GetUser(expect.ID)
	require.ErrorIs(t, err, storages.ErrUserNotFound)
}

func TestUserServiceListUsers(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	record1, err := service.CreateUser("Test", "1@1.com", "admin")
	require.NoError(t, err)

	record2, err := service.CreateUser("Test", "2@2.com", "user")
	require.NoError(t, err)

	require.Equal(t, []*user.User{record1, record2}, service.ListUsers())
}

func TestUserServiceListUsersWithRole(t *testing.T) {
	t.Parallel()

	service := services.NewUserService(memory.NewInMemoryUserRepo())

	record1, err := service.CreateUser("Test", "1@1.com", "admin")
	require.NoError(t, err)

	record2, err := service.CreateUser("Test", "2@2.com", "user")
	require.NoError(t, err)

	require.Equal(t, []*user.User{record1}, service.ListUsersWithRole("admin"))
	require.Equal(t, []*user.User{record2}, service.ListUsersWithRole("user"))
}
