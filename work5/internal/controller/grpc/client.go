package usergrpc

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mch735/education/work5/internal/entity"
	pb "github.com/mch735/education/work5/pkg/proto/gen/userpb"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.UserServiceClient
}

func NewClient(host string, port int) (*Client, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	service := pb.NewUserServiceClient(conn)

	return &Client{conn: conn, service: service}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) GetUsers(ctx context.Context) ([]*entity.User, error) {
	result, err := c.service.GetUsers(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("c.service.GetUsers: %w", err)
	}

	data := result.GetUsers()

	users := make([]*entity.User, 0, len(data))
	for _, user := range data {
		users = append(users, c.toEntityUser(user))
	}

	return users, nil
}

func (c *Client) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	result, err := c.service.GetUserByID(ctx, &pb.ID{Id: id})
	if err != nil {
		return nil, fmt.Errorf("c.service.GetUserByID: %w", err)
	}

	return c.toEntityUser(result), nil
}

func (c *Client) Create(ctx context.Context, name, email, role string) (*entity.User, error) {
	result, err := c.service.Create(ctx, &pb.UserRequest{Name: name, Email: email, Role: role})
	if err != nil {
		return nil, fmt.Errorf("c.service.Create: %w", err)
	}

	return c.toEntityUser(result), nil
}

func (c *Client) Update(ctx context.Context, id, name, email, role string) (*entity.User, error) {
	result, err := c.service.Create(ctx, &pb.UserRequest{Id: id, Name: name, Email: email, Role: role})
	if err != nil {
		return nil, fmt.Errorf("c.service.Update: %w", err)
	}

	return c.toEntityUser(result), nil
}

func (c *Client) Delete(ctx context.Context, id string) error {
	_, err := c.service.Delete(ctx, &pb.ID{Id: id})
	if err != nil {
		return fmt.Errorf("c.service.Delete: %w", err)
	}

	return nil
}

func (c *Client) toEntityUser(user *pb.UserResponse) *entity.User {
	return &entity.User{
		ID:        user.GetId(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		Role:      user.GetRole(),
		CreatedAt: user.GetCreatedAt().AsTime(),
		UpdatedAt: user.GetUpdatedAt().AsTime(),
	}
}
