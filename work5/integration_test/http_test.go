//go:build integration

package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/entity/dto"
	"github.com/stretchr/testify/require"
)

const baseURL = "http://localhost:8080"

// POST /v1/users
func TestHttpUserCreate(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	result, err := create(ctx, "User", "1@1.com", "admin")
	require.NoError(t, err)
	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)
}

// GET /v1/users
func TestHttpUserList(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	create(ctx, "User", "1@1.com", "admin")

	resp, err := do(ctx, http.MethodGet, "/v1/users", http.NoBody)
	require.NoError(t, err)

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var result []*entity.User

	err = json.Unmarshal(data, &result)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

// GET /v1/users/{id}
func TestHttpUserShow(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user, err := create(ctx, "User", "1@1.com", "admin")
	require.NoError(t, err)

	resp, err := do(ctx, http.MethodGet, "/v1/users/"+user.ID, http.NoBody)
	require.NoError(t, err)

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var result *entity.User

	err = json.Unmarshal(data, &result)
	require.NoError(t, err)
	require.Equal(t, user.ID, result.ID)
	require.Equal(t, user.Name, result.Name)
	require.Equal(t, user.Role, result.Role)
}

// PUT /v1/users/{id}
func TestHttpUserUpdate(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user, err := create(ctx, "User", "1@1.com", "admin")
	require.NoError(t, err)

	u := dto.User{Name: "User updated", Email: "Q@Q.com", Role: "user"}

	data, err := json.Marshal(u)
	require.NoError(t, err)

	resp, err := do(ctx, http.MethodPut, "/v1/users/"+user.ID, bytes.NewBuffer(data))
	require.NoError(t, err)

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	require.NoError(t, err)

	var result *entity.User

	err = json.Unmarshal(data, &result)
	require.NoError(t, err)
	require.Equal(t, user.ID, result.ID)
	require.Equal(t, "User updated", result.Name)
	require.Equal(t, "Q@Q.com", result.Email)
	require.Equal(t, "user", result.Role)
}

// DELETE /v1/users/{id}
func TestHttpUserDelete(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user, err := create(ctx, "User", "1@1.com", "admin")
	require.NoError(t, err)

	resp, err := do(ctx, http.MethodDelete, "/v1/users/"+user.ID, http.NoBody)
	require.NoError(t, err)

	defer resp.Body.Close()
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func do(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	url, err := url.JoinPath(baseURL, path)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.DefaultClient.Do: %w", err)
	}

	return result, nil
}

func create(ctx context.Context, name, email, role string) (*entity.User, error) {
	u := dto.User{Name: name, Email: email, Role: role}

	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	resp, err := do(ctx, http.MethodPost, "/v1/users", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	var user *entity.User

	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return user, nil
}
