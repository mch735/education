package usergrpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/mch735/education/work5/config"
	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/usecase"
	pb "github.com/mch735/education/work5/pkg/proto/gen/userpb"
)

type Server struct {
	conf *config.GRPC
	grpc *grpc.Server
}

func NewServer(conf *config.GRPC) *Server {
	server := grpc.NewServer()

	return &Server{conf: conf, grpc: server}
}

func (s *Server) RegisterServiceServer(service pb.UserServiceServer) {
	pb.RegisterUserServiceServer(s.grpc, service)
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.conf.Port))
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	err = s.grpc.Serve(listener)
	if err != nil {
		return fmt.Errorf("s.grpc.Serve: %w", err)
	}

	return nil
}

func (s *Server) Shutdown() error {
	s.grpc.GracefulStop()

	return nil
}

func NewUserServiceServer(uc usecase.User) *ImplementedUserServiceServer {
	return &ImplementedUserServiceServer{uc: uc}
}

type ImplementedUserServiceServer struct {
	pb.UnimplementedUserServiceServer
	uc usecase.User
}

func (s *ImplementedUserServiceServer) GetUsers(ctx context.Context, _ *empty.Empty) (*pb.UsersResponse, error) {
	users, err := s.uc.GetUsers(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("s.uc.GetUsers: %v", err))

		return nil, status.Errorf(codes.Internal, "method GetUsers error")
	}

	result := make([]*pb.UserResponse, 0, len(users))
	for _, user := range users {
		result = append(result, s.toUserResponse(user))
	}

	return &pb.UsersResponse{Users: result}, nil
}

func (s *ImplementedUserServiceServer) GetUserByID(ctx context.Context, id *pb.ID) (*pb.UserResponse, error) {
	user, err := s.uc.GetUserByID(ctx, id.GetId())
	if err != nil {
		slog.Error(fmt.Sprintf("s.uc.GetUserByID: %v", err))

		return nil, status.Errorf(codes.Internal, "method GetUserByID error")
	}

	return s.toUserResponse(user), nil
}

func (s *ImplementedUserServiceServer) Create(ctx context.Context, ur *pb.UserRequest) (*pb.UserResponse, error) {
	user := entity.User{
		Name:  ur.GetName(),
		Email: ur.GetEmail(),
		Role:  ur.GetRole(),
	}

	err := s.uc.Create(ctx, &user)
	if err != nil {
		slog.Error(fmt.Sprintf("s.uc.Create: %v", err))

		return nil, status.Errorf(codes.Internal, "method Create error")
	}

	return s.toUserResponse(&user), nil
}

func (s *ImplementedUserServiceServer) Update(ctx context.Context, ur *pb.UserRequest) (*pb.UserResponse, error) {
	user := entity.User{
		ID:    ur.GetId(),
		Name:  ur.GetName(),
		Email: ur.GetEmail(),
		Role:  ur.GetRole(),
	}

	err := s.uc.Update(ctx, &user)
	if err != nil {
		slog.Error(fmt.Sprintf("s.uc.Update: %v", err))

		return nil, status.Errorf(codes.Internal, "method Update error")
	}

	return s.toUserResponse(&user), nil
}

func (s *ImplementedUserServiceServer) Delete(ctx context.Context, id *pb.ID) (*empty.Empty, error) {
	err := s.uc.Delete(ctx, id.GetId())
	if err != nil {
		slog.Error(fmt.Sprintf("s.uc.Delete: %v", err))

		return nil, status.Errorf(codes.Internal, "method Delete error")
	}

	return &empty.Empty{}, nil
}

func (s *ImplementedUserServiceServer) toUserResponse(user *entity.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
