package handlers

import (
	"context"

	pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"
)

func (u *UserHandler) UserFoodByName(ctx context.Context, p *pb.FoodByName) (*pb.MenuItem, error) {
	result, err := u.svc.UserFoodByName(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserHandler) UserMenuByID(ctx context.Context, p *pb.MenuID) (*pb.MenuItem, error) {
	result, err := u.svc.UserMenuID(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserHandler) UserMenuList(ctx context.Context, p *pb.RNoparam) (*pb.MenuList, error) {
	result, err := u.svc.UserMenuList(p)
	if err != nil {
		return nil, err
	}

	return result, nil
}
