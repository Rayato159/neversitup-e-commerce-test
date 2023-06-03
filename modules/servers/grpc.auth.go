package servers

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authProto"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authUsecase"
	"google.golang.org/grpc"
)

func (s *server) StartAuthGRPCServer() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.App().Host(), s.cfg.App().Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	modules := NewModule(s, nil, grpcServer)
	pb.RegisterAuthServicesServer(
		grpcServer,
		NewAuthGRPCModule(modules.NewAuthModule().Usecase()),
	)

	log.Printf("auth grpc server is starting on %v", s.cfg.App().Url())
	grpcServer.Serve(lis)
}

type authGrpcServer struct {
	authUsecase authUsecase.IAuthUsecase
	pb.UnimplementedAuthServicesServer
}

// gRPC
func NewAuthGRPCModule(authUsecase authUsecase.IAuthUsecase) *authGrpcServer {
	return &authGrpcServer{
		authUsecase: authUsecase,
	}
}

func (h *authGrpcServer) FindAccessToken(ctx context.Context, in *pb.FindAccessTokenReq) (*pb.FindAccessTokenRes, error) {
	return &pb.FindAccessTokenRes{
		Result: h.authUsecase.FindAccessToken(in.UserId, in.AccessToken),
	}, nil
}
