package user_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mch735/education/work2/internal/storages"
	"github.com/mch735/education/work2/internal/storages/memory"
	"github.com/mch735/education/work2/internal/storages/mock"
	"github.com/mch735/education/work2/internal/user"
)

func TestUserServiceValidateError(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	_, err := service.CreateUser("Test", "1@1.com", "qwe")
	require.ErrorIs(t, err, user.ErrInvalidRole)

	_, err = service.CreateUser("Test", "example.com", "admin")
	require.ErrorIs(t, err, user.ErrInvalidEmail)

	_, err = service.CreateUser("", "1@1.com", "user")
	require.ErrorIs(t, err, user.ErrInvalidName)
}

func TestUserServiceSaveError(t *testing.T) {
	t.Parallel()

	repo := mock.NewErrorUserRepo()

	_, err := user.NewService(repo).CreateUser("Test", "1@1.com", "admin")
	require.ErrorIs(t, err, storages.ErrUserExist)
}

func TestUserServiceSave(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	record, _ := service.CreateUser("Test", "1@1.com", "user")
	require.Equal(t, "Test", record.Name)
	require.Equal(t, "1@1.com", record.Email)
	require.Equal(t, "user", record.Role)
}

func TestUserServiceGetUserError(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	_, err := service.GetUser("10")
	require.ErrorIs(t, err, storages.ErrUserNotFound)
}

func TestUserServiceGetUser(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	expect, err := service.CreateUser("Test", "1@1.com", "user")
	require.NoError(t, err)

	actual, err := service.GetUser(expect.ID)
	require.NoError(t, err)
	require.Equal(t, expect, actual)
}

func TestUserServiceRemoveUserError(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	err := service.RemoveUser("10")
	require.ErrorIs(t, err, storages.ErrUserNotFound)
}

func TestUserServiceRemoveUser(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	expect, err := service.CreateUser("Test", "1@1.com", "user")
	require.NoError(t, err)

	require.NoError(t, service.RemoveUser(expect.ID))

	_, err = service.GetUser(expect.ID)
	require.ErrorIs(t, err, storages.ErrUserNotFound)
}

func TestUserServiceListUsers(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	record1, err := service.CreateUser("Test", "1@1.com", "admin")
	require.NoError(t, err)

	record2, err := service.CreateUser("Test", "2@2.com", "user")
	require.NoError(t, err)

	require.Equal(t, []*user.User{record1, record2}, service.ListUsers())
}

func TestUserServiceListUsersWithRole(t *testing.T) {
	t.Parallel()

	service := user.NewService(memory.NewUserRepo())

	record1, err := service.CreateUser("Test", "1@1.com", "admin")
	require.NoError(t, err)

	record2, err := service.CreateUser("Test", "2@2.com", "user")
	require.NoError(t, err)

	require.Equal(t, []*user.User{record1}, service.ListUsersWithRole("admin"))
	require.Equal(t, []*user.User{record2}, service.ListUsersWithRole("user"))
}
