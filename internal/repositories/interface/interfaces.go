package interfaces

import "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/model"

// UserRepository defines the methods for user repository
type UserRepository interface {
	CreateUser(user *model.UserModel) error
	FindUserByPhone(phone string) (*model.UserModel, error)
}
