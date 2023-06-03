package servers

import (
	"context"
	"time"

	pb "github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersProto"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersUsecase"
)

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
