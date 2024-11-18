package handler

import (
	"context"
	"github.com/tittuvarghese/ss-go-core/logger"
	"github.com/tittuvarghese/ss-go-customer-service/core/database"
	"github.com/tittuvarghese/ss-go-customer-service/models"
	"github.com/tittuvarghese/ss-go-customer-service/proto"
	"github.com/tittuvarghese/ss-go-customer-service/service"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
	return &Server{GrpcServer: grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))}
}

func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", ":"+port)
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
	user.Type = req.Type

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

	user, err := service.GetProfile(req.Userid, s.RdbInstance)

	if err != nil {
		return &proto.GetProfileResponse{}, err
	}

	return &proto.GetProfileResponse{
		Userid:    user.ID.String(),
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Type:      user.Type,
	}, nil
}
