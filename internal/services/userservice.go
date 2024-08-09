package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ratheeshkumar/restaurant_user_serviceV1/config"
	menupb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/menus/pb"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/model"
	pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"
	userrepo "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/repositories/interface"
	services_inter "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/services/interfaces"
	utils "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/utils"
)

type UserService struct {
	repo         userrepo.UserRepository
	client       menupb.MenuServiceClient
	redisClient  *config.RedisService
	twilioClient *config.TwilioService
}

// Login implements services_inter.UserService.
func (u *UserService) Login(userpb *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := u.repo.FindUserByPhone(userpb.Phone)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Generate the JWT token with the user's phone, user ID, and role
	token, err := utils.GenerateToken(user.Phone, uint(user.UserID))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// token, err := utils.GenerateToken(user.Phone, uint(user.UserID,))
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to generate token: %w", err)
	// }

	return &pb.LoginResponse{Token: token}, nil
}

// Signup implements services_inter.UserService.
func (u *UserService) Signup(userpb *pb.SignupRequest) (*pb.SignupRespnse, error) {
	// Check if the user already exists in the database
	existingUser, err := u.repo.FindUserByPhone(userpb.Phone)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	// Send OTP using Twilio
	_, err = u.twilioClient.SendTwilioOTP(userpb.Phone)
	if err != nil {
		return nil, err
	}

	// Create a user model
	user := &model.UserModel{Phone: userpb.Phone}

	// Create a Redis key with the user's phone number
	key := fmt.Sprintf("user:%s", user.Phone)
	fmt.Println("sETKEY", key)
	// Marshal the user data to JSON
	userData, err := json.Marshal(&user)
	log.Println("marshal", userData)
	if err != nil {
		return nil, err
	}

	// Store user data in Redis with a TTL of 5 minutes
	err = u.redisClient.SetDataInRedis(key, userData, time.Minute*5)
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully set data in Redis for key: %s", key)
	// Return a response indicating that the signup has been initiated
	return &pb.SignupRespnse{Message: "Signup initiated, please verify OTP."}, nil
}

// // VerifyOTP implements services_inter.UserService.
func (u *UserService) VerifyOTP(userpb *pb.VerifyOTPRequest) (*pb.VerifyOTPRespnse, error) {
	// phone := userpb.Phone
	// otp := userpb.Otp

	err := u.twilioClient.VerifyTwilioOTP(userpb.Phone, userpb.Otp)
	if err != nil {
		log.Fatal("unable verify the otp")
	}
	fmt.Println("userpb")
	key := fmt.Sprintf("user:%s", userpb.Phone)
	fmt.Println("GETKEY", key)
	// userData, err := u.redisClient.GetFromRedis(key).Result()
	userData, err := u.redisClient.GetFromRedis(key)
	if err != nil {
		return nil, err
	}

	var user model.UserModel
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}
	err = u.repo.CreateUser(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := utils.GenerateToken(userpb.Phone, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &pb.VerifyOTPRespnse{Token: token}, nil
}

// NewUserService creates a new instance of UserService

func NewUserService(repo userrepo.UserRepository, client menupb.MenuServiceClient, redisClient *config.RedisService, twilioClient *config.TwilioService) services_inter.UserService {
	return &UserService{
		repo:         repo,
		client:       client,
		redisClient:  redisClient,
		twilioClient: twilioClient,
	}
}
