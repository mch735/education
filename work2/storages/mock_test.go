package storages_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mch735/education/work2/models/user"
	"github.com/mch735/education/work2/storages"
)

func TestMockRepoSaveError(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockErrorUserRepo()

	record := user.User{ID: "10", Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.ErrorIs(t, storages.ErrUserExist, repo.Save(record))
}

func TestMockRepoSave(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockSuccessUserRepo()

	record := user.User{ID: "10", Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record))
}

func TestMockRepoFindByIDError(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockErrorUserRepo()

	_, err := repo.FindByID("10")
	require.ErrorIs(t, storages.ErrUserNotFound, err)
}

func TestMockRepoFindByID(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockSuccessUserRepo()

	_, err := repo.FindByID("10")
	require.NoError(t, err)
}

func TestMockRepoDeleteByIDError(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockErrorUserRepo()
	require.ErrorIs(t, storages.ErrUserNotFound, repo.DeleteByID("10"))
}

func TestMockRepoDeleteByID(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockSuccessUserRepo()
	require.NoError(t, repo.DeleteByID("10"))
}

func TestMockRepoFindAll(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockSuccessUserRepo()
	require.Equal(t, []user.User{}, repo.FindAll())
}

func TestMockRepoFilterFunc(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockSuccessUserRepo()
	require.Equal(t, []user.User{}, repo.FilterFunc(func(user user.User) bool {
		return user.Role == "user"
	}))
}
