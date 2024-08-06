package handlers

import (
	"context"
	"fmt"

	pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"
	services_inter "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/services/interfaces"
)

type UserHandler struct {
	pb.UserServicesServer
	svc services_inter.UserService
}

func NewUserHandler(svc services_inter.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}
func (u *UserHandler) UseSignup(ctx context.Context, p *pb.SignupRequest) *pb.SignupRespnse {
	result, err := u.svc.Signup(p)

	if err != nil {
		return &pb.SignupRespnse{
			Message: "Failed to send the Otp",
		}
	}
	msg := fmt.Sprintf("Otp send for the verification %v", result)
	return &pb.SignupRespnse{
		Message: msg,
	}

}

func (u *UserHandler) VerifyOTP(ctx context.Context, p *pb.VerifyOTPRequest) (*pb.VerifyOTPRespnse, error) {
	result, err := u.svc.VerifyOTP(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserHandler) Login(ctx context.Context, p *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := u.svc.Login(p)
	if err != nil {
		return nil, err
	}

	return result, nil

}
