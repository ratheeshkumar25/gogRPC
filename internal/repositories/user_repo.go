package repositories

import (
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/model"
	userrepo "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/repositories/interface"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

// NewUserRepo creates an instance of user repo
func NewUserRepo(db *gorm.DB) userrepo.UserRepository {
	return &UserRepo{DB: db}
}

// CreateUser creates a new user in the database, else returns an error
func (u *UserRepo) CreateUser(user *model.UserModel) error {
	if err := u.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindUserByPhone finds a user by phone number, else returns an error
func (u *UserRepo) FindUserByPhone(phone string) (*model.UserModel, error) {
	var user model.UserModel
	err := u.DB.Where("phone = ?", phone).First(&user).Error
	return &user, err
}
