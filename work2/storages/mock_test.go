package storages_test

import (
	"testing"
	"time"

	"github.com/mch735/education/work2/models/user"
	"github.com/mch735/education/work2/storages"
	"github.com/stretchr/testify/require"
)

func TestMockRepoSave(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockUserRepo()

	record := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record))
}

func TestMockRepoFindByID(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockUserRepo()

	_, err := repo.FindByID(10)
	require.NoError(t, err)
}

func TestMockRepoFindAll(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockUserRepo()
	require.Equal(t, []user.User{}, repo.FindAll())
}

func TestMockRepoDeleteByID(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockUserRepo()
	require.NoError(t, repo.DeleteByID(10))
}

func TestMockRepoFilterFunc(t *testing.T) {
	t.Parallel()

	repo := storages.NewMockUserRepo()
	require.Equal(t, []user.User{}, repo.FilterFunc(func(user user.User) bool {
		return user.Role == "user"
	}))
}
