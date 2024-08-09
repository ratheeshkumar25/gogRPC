package services_inter

import pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"

type UserService interface {
	Signup(userpb *pb.SignupRequest) (*pb.SignupRespnse, error)
	VerifyOTP(userpb *pb.VerifyOTPRequest) (*pb.VerifyOTPRespnse, error)
	Login(userpb *pb.LoginRequest) (*pb.LoginResponse, error)
	UserFoodByName(*pb.FoodByName) (*pb.MenuItem, error)
	UserMenuByID(*pb.MenuID) (*pb.MenuItem, error)
	UserMenuList(*pb.RNoParam) (*pb.MenuList, error)
}
