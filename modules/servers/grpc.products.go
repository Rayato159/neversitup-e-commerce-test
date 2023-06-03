package servers

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsProto"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsUsecase"
	"google.golang.org/grpc"
)

func (s *server) StartProductsGRPCServer() {
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.App().Host(), s.cfg.App().Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	modules := NewModule(s, nil, grpcServer)
	pb.RegisterProductsServicesServer(
		grpcServer,
		NewProductsGRPCModule(modules.NewProductsModule().Usecase()),
	)

	log.Printf("products grpc server is starting on %v", s.cfg.App().Url())
	grpcServer.Serve(lis)
}

type productsGrpcServer struct {
	productsUsecase productsUsecase.IProductsUsecase
	pb.UnimplementedProductsServicesServer
}

// gRPC
func NewProductsGRPCModule(productsUsecase productsUsecase.IProductsUsecase) *productsGrpcServer {
	return &productsGrpcServer{
		productsUsecase: productsUsecase,
	}
}
