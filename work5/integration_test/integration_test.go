//go:build integration

package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/entity/dto"
)

const baseURL = "http://localhost:8080"

var userList []*entity.User

// POST /v1/users
func TestUserCreate(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user := dto.User{Name: "User1", Email: "1@1.com", Role: "admin"}
	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("TestUserCreate: json.Marshal: %v", err)
	}

	resp, err := doRequestWithContext(ctx, http.MethodPost, "/v1/users", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("TestUserCreate: doRequestWithContext: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("TestUserCreate: resp.StatusCode: %d", resp.StatusCode)
	}
}

// GET /v1/users
func TestUserList(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	resp, err := doRequestWithContext(ctx, http.MethodGet, "/v1/users", http.NoBody)
	if err != nil {
		t.Fatalf("TestUserList: doRequestWithContext: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("TestUserList: resp.StatusCode: %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("TestUserList: ioutil.ReadAll: %v", err)
	}

	err = json.Unmarshal(data, &userList)
	if err != nil {
		t.Fatalf("TestUserList: json.Unmarshal: %v", err)
	}
}

// GET /v1/users/{id}
func TestUserShow(t *testing.T) {
	ID := userList[0].ID

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	resp, err := doRequestWithContext(ctx, http.MethodGet, "/v1/users/"+ID, http.NoBody)
	if err != nil {
		t.Fatalf("TestUserShow: doRequestWithContext: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("TestUserShow: resp.StatusCode: %d", resp.StatusCode)
	}
}

// PUT /v1/users/{id}
func TestUserUpdate(t *testing.T) {
	ID := userList[0].ID

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user := dto.User{Name: "User Updated", Email: "Q@Q.com", Role: "user"}
	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("TestUserCreate: json.Marshal: %v", err)
	}

	resp, err := doRequestWithContext(ctx, http.MethodPut, "/v1/users/"+ID, bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("TestUserUpdate: doRequestWithContext: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("TestUserUpdate: resp.StatusCode: %d", resp.StatusCode)
	}
}

// DELETE /v1/users/{id}
func TestUserDelete(t *testing.T) {
	ID := userList[0].ID

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	resp, err := doRequestWithContext(ctx, http.MethodDelete, "/v1/users/"+ID, http.NoBody)
	if err != nil {
		t.Fatalf("TestUserDelete: doRequestWithContext: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("TestUserDelete: resp.StatusCode: %d", resp.StatusCode)
	}
}

func doRequestWithContext(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
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
