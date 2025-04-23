package web

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/entity/dto"
)

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

func (c *Client) GetUsers(ctx context.Context) ([]*entity.User, error) {
	resp, err := c.RequestWithContext(ctx, http.MethodGet, "/v1/users", http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	var users []*entity.User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return users, nil
}

func (c *Client) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	resp, err := c.RequestWithContext(ctx, http.MethodGet, "/v1/users/"+id, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
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

func (c *Client) Create(ctx context.Context, name, email, role string) (*entity.User, error) {
	u := dto.User{Name: name, Email: email, Role: role}

	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	resp, err := c.RequestWithContext(ctx, http.MethodPost, "/v1/users", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("c.RequestWithContext: %w", err)
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

func (c *Client) Update(ctx context.Context, id, name, email, role string) (*entity.User, error) {
	u := dto.User{Name: name, Email: email, Role: role}

	data, err := json.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	resp, err := c.RequestWithContext(ctx, http.MethodPut, "/v1/users/"+id, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	var user *entity.User

	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return user, nil
}

func (c *Client) Delete(ctx context.Context, id string) error {
	resp, err := c.RequestWithContext(ctx, http.MethodDelete, "/v1/users/"+id, http.NoBody)
	if err != nil {
		return fmt.Errorf("c.RequestWithContext: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

func (c *Client) RequestWithContext(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	url, err := url.JoinPath(c.baseURL, path)
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
