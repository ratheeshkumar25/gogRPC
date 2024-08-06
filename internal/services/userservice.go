package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	menupb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/menus/pb"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/model"
	pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"
	userrepo "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/repositories/interface"
	services_inter "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/services/interfaces"
	utils "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/utils"
	"github.com/redis/go-redis/v9"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type UserService struct {
	repo         userrepo.UserRepository
	client       menupb.MenuServiceClient
	redisClient  *redis.Client
	twilioClient *twilio.RestClient
}

// Login implements services_inter.UserService.
func (u *UserService) Login(userpb *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := u.repo.FindUserByPhone(userpb.Phone)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = u.sendOTP(userpb.Phone)
	if err != nil {
		return nil, fmt.Errorf("failed to send OTP: %w", err)
	}

	token, err := utils.GenerateToken(user.Phone, uint(user.UserID))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.LoginResponse{Token: token}, nil
}

// Signup implements services_inter.UserService.
func (u *UserService) Signup(userpb *pb.SignupRequest) (*pb.SignupRespnse, error) {
	existingUser, err := u.repo.FindUserByPhone(userpb.Phone)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	err = u.sendOTP(userpb.Phone)
	if err != nil {
		return nil, err
	}

	user := &model.UserModel{Phone: userpb.Phone}
	key := fmt.Sprintf("user:%s", user.Phone)
	userData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = u.redisClient.Set(context.Background(), key, userData, time.Minute*5).Err()
	if err != nil {
		return nil, err
	}

	return &pb.SignupRespnse{Message: "Signup initiated, please verify OTP."}, nil
}

// VerifyOTP implements services_inter.UserService.
func (u *UserService) VerifyOTP(userpb *pb.VerifyOTPRequest) (*pb.VerifyOTPRespnse, error) {
	params := verify.CreateVerificationCheckParams{}
	params.SetTo(userpb.Phone)
	params.SetCode(userpb.Otp)

	response, err := u.twilioClient.VerifyV2.CreateVerificationCheck("YOUR_SERVICE_SID", &params)
	if err != nil || response.Status == nil || *response.Status != "approved" {
		return nil, errors.New("invalid OTP")
	}

	key := fmt.Sprintf("user:%s", userpb.Phone)
	userData, err := u.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		var user model.UserModel
		err = json.Unmarshal([]byte(userData), &user)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
		}
		err = u.repo.CreateUser(&user)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	token, err := utils.GenerateToken(userpb.Phone, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.VerifyOTPRespnse{Token: token}, nil
}

func (u *UserService) sendOTP(phone string) error {
	params := verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	_, err := u.twilioClient.VerifyV2.CreateVerification("YOUR_SERVICE_SID", &params)
	return err
}

// NewUserService creates a new instance of UserService
func NewUserService(repo userrepo.UserRepository, client menupb.MenuServiceClient, redisClient *redis.Client, twilioClient *twilio.RestClient) services_inter.UserService {
	return &UserService{
		repo:         repo,
		client:       client,
		redisClient:  redisClient,
		twilioClient: twilioClient,
	}
}
