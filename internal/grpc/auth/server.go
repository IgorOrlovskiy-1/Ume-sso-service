package auth

import (
	ssov1 "github.com/IgorOrlovskiy-1/Ume-protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"context"
)

const emptyValue = 0

type Auth interface {
	Login(ctx context.Context,
	username string,
	password string,
	appId int,
	) (token string, err error)
	RegisterNewUser(ctx context.Context,
	username string,
	password string,
	) (userId int64, err error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, req.GetUsername(), req.GetPassword(), req.GetAppId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ssov1.LoginRequest{
		Token: token,
	}, nil

}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	if err := validateRegister(req); err != nil {
		return nil, err
	}

	userId, err := s.auth.RegisterNewUser(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		//TODO: refactoring error
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &ssov1.RegisterRequest{
		UserId: userId,
	}, nil
}

func validateLogin(req *ssov1.LoginRequest) (string, error) {
	if req.GetUsername() == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app_id is required") 
	}
}

func validateRegister(req *ssov1.LoginRequest) (string, error) {
	if req.GetUsername() == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
}