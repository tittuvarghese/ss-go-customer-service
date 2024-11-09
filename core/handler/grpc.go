package handler

import (
	"context"
	"github.com/tittuvarghese/core/logger"
	"github.com/tittuvarghese/customer-service/core/database"
	"github.com/tittuvarghese/customer-service/models"
	"github.com/tittuvarghese/customer-service/proto"
	"github.com/tittuvarghese/customer-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	proto.UnimplementedAuthServiceServer
	GrpcServer  *grpc.Server
	RdbInstance *database.RelationalDatabase
}

var log = logger.NewLogger("customer-service")

func NewGrpcServer() *Server {
	return &Server{GrpcServer: grpc.NewServer()}
}

func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Error("Failed to listen", err)
	}

	proto.RegisterAuthServiceServer(s.GrpcServer, s)

	// Register reflection service on gRPC server
	reflection.Register(s.GrpcServer)
	log.Info("GRPC server is listening on port " + port)
	if err := s.GrpcServer.Serve(lis); err != nil {
		log.Error("failed to serve", err)
	}
}

func (s *Server) mustEmbedUnimplementedAuthServiceServer() {
	log.Error("implement me", nil)
}

// Register a new user
func (s *Server) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {

	var user models.User
	user.Username = req.Username
	user.Password = req.Password
	user.Firstname = req.Firstname
	user.Lastname = req.Lastname

	err := service.CreateUser(user, s.RdbInstance)
	if err != nil {
		return &proto.RegisterResponse{
			Message: "Failed to register the user. error: " + err.Error(),
		}, err
	}

	return &proto.RegisterResponse{
		Message: "User registered successfully",
	}, nil
}

// Login
func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	var request models.LoginRequest
	request.Username = req.Username
	request.Password = req.Password

	token, err := service.AuthenticateUser(request, s.RdbInstance)

	if err != nil {
		return &proto.LoginResponse{
			Token: "Unable to authenticate the user",
		}, err
	}

	return &proto.LoginResponse{
		Token: token,
	}, nil
}

// GetProfile
func (s *Server) GetProfile(ctx context.Context, req *proto.GetProfileRequest) (*proto.GetProfileResponse, error) {

	return &proto.GetProfileResponse{
		Userid:    "1",
		Username:  "john_doe",
		Firstname: "Profile fetched successfully",
		Lastname:  "Profile fetched successfully",
	}, nil
}
