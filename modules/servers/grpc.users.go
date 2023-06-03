package servers

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersProto"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersUsecase"
	"google.golang.org/grpc"
)

func (s *server) StartUsersGRPCServer() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.App().Host(), s.cfg.App().Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	modules := NewModule(s, nil, grpcServer)
	pb.RegisterUsersServicesServer(
		grpcServer,
		NewUsersGRPCModule(modules.NewUsersModule().Usecase()),
	)

	log.Printf("users grpc server is starting on %v", s.cfg.App().Url())
	grpcServer.Serve(lis)
}

type usersGrpcServer struct {
	usersUsecase usersUsecase.IUsersUsecase
	pb.UnimplementedUsersServicesServer
}

// gRPC
func NewUsersGRPCModule(usersUsecase usersUsecase.IUsersUsecase) *usersGrpcServer {
	return &usersGrpcServer{
		usersUsecase: usersUsecase,
	}
}

func (h *usersGrpcServer) FindOneUserByUsername(ctx context.Context, in *pb.FindOneUserByUsernameReq) (*pb.UserForAll, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	user, err := h.usersUsecase.FindOneUserByUsername(in.Username)
	if err != nil {
		return nil, err
	}

	return &pb.UserForAll{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (h *usersGrpcServer) FindOneUserForAllById(ctx context.Context, in *pb.FindOneUserByIdReq) (*pb.UserForAll, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	user, err := h.usersUsecase.FindOneUserForAllById(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UserForAll{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (h *usersGrpcServer) FindOneUserById(ctx context.Context, in *pb.FindOneUserByIdReq) (*pb.FindOneUserByIdRes, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	isUser := h.usersUsecase.FindOneUserById(in.Id)

	return &pb.FindOneUserByIdRes{
		Result: isUser,
	}, nil
}
