package storages_test

import (
	"testing"
	"time"

	"github.com/mch735/education/work2/models/user"
	"github.com/mch735/education/work2/storages"

	"github.com/stretchr/testify/require"
)

func TestInMemoryRepoSaveError(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record1 := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record1))

	record2 := user.User{ID: 10, Name: "Test2", Email: "2@2.com", Role: "user", CreatedAt: time.Now()}
	require.ErrorIs(t, repo.Save(record2), storages.ErrUserExist)
}

func TestInMemoryRepoSave(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record))
}

func TestInMemoryRepoFindByIDError(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record))

	_, err := repo.FindByID(20)
	require.ErrorIs(t, err, storages.ErrUserNotFound)
}

func TestInMemoryRepoFindByID(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record))

	result, _ := repo.FindByID(10)
	require.Equal(t, record, result)
}

func TestInMemoryRepoFindAll(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record1 := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record1))

	record2 := user.User{ID: 20, Name: "Test2", Email: "2@2.com", Role: "user", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record2))

	result := repo.FindAll()
	require.Equal(t, []user.User{record1, record2}, result)
}

func TestInMemoryRepoDeleteByIDError(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record))

	require.ErrorIs(t, repo.DeleteByID(20), storages.ErrUserNotFound)
}

func TestInMemoryRepoDeleteByID(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record := user.User{ID: 10, Name: "Test1", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record))

	require.NoError(t, repo.DeleteByID(10))
}

func TestInMemoryRepoFilterFunc(t *testing.T) {
	t.Parallel()

	repo := storages.NewInMemoryUserRepo()

	record1 := user.User{ID: 10, Name: "Test", Email: "1@1.com", Role: "admin", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record1))

	record2 := user.User{ID: 20, Name: "Test", Email: "2@2.com", Role: "user", CreatedAt: time.Now()}
	require.NoError(t, repo.Save(record2))

	require.Equal(t, []user.User{record2}, repo.FilterFunc(func(user user.User) bool {
		return user.Role == "user"
	}))

	require.Equal(t, []user.User{record1, record2}, repo.FilterFunc(func(user user.User) bool {
		return user.Name == "Test"
	}))
}
