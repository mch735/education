//go:build integration

package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mch735/education/work5/pkg/proto/gen/userpb"
)

const (
	host = "localhost"
	port = 50051
)

func TestGrpcUserCreate(t *testing.T) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	defer conn.Close()
	service := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	result, err := service.Create(ctx, &pb.UserRequest{Name: "Test", Email: "1@1.com", Role: "admin"})
	require.NoError(t, err)
	require.Equal(t, "Test", result.GetName())
	require.Equal(t, "1@1.com", result.GetEmail())
	require.Equal(t, "admin", result.GetRole())
}

func TestGrpcUserList(t *testing.T) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	defer conn.Close()
	service := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	service.Create(ctx, &pb.UserRequest{Name: "Test", Email: "1@1.com", Role: "admin"})

	result, err := service.GetUsers(ctx, &empty.Empty{})
	require.NoError(t, err)
	require.NotEmpty(t, result.GetUsers())
}

func TestGrpcUserShow(t *testing.T) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	defer conn.Close()
	service := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user, err := service.Create(ctx, &pb.UserRequest{Name: "Test", Email: "1@1.com", Role: "admin"})
	result, err := service.GetUserByID(ctx, &pb.ID{Id: user.GetId()})

	require.NoError(t, err)
	require.NotZero(t, result.GetId())
	require.NotZero(t, result.GetCreatedAt().AsTime())
}

func TestGrpcUserUpdate(t *testing.T) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	defer conn.Close()
	service := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user, err := service.Create(ctx, &pb.UserRequest{Name: "Test", Email: "1@1.com", Role: "admin"})
	result, err := service.Create(ctx, &pb.UserRequest{Id: user.GetId(), Name: "Update", Email: "2@2.com", Role: "guest"})

	require.NoError(t, err)
	require.Equal(t, "Update", result.GetName())
	require.Equal(t, "2@2.com", result.GetEmail())
	require.Equal(t, "guest", result.GetRole())
}

func TestGrpcUserDelete(t *testing.T) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	defer conn.Close()
	service := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Second)
	defer cancel()

	user, err := service.Create(ctx, &pb.UserRequest{Name: "Test", Email: "1@1.com", Role: "admin"})
	result, err := service.Delete(ctx, &pb.ID{Id: user.GetId()})

	require.NoError(t, err)
	require.IsType(t, &empty.Empty{}, result)
}
