package services

import (
	menu "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/menus/handlers"
	pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"
)

// UserFoodByName implements services_inter.UserService.
func (u *UserService) UserFoodByName(userpb *pb.FoodByName) (*pb.MenuItem, error) {
	result, err := menu.FetchMenuBYNameHandler(u.client, userpb)
	if err != nil {
		return nil, err
	}

	return &pb.MenuItem{
		Id:        result.Id,
		Category:  result.Category,
		Name:      result.Name,
		Price:     result.Price,
		Foodimage: result.Foodimage,
		Duration:  result.Duration,
	}, nil

}

// UserMenuID implements services_inter.UserService.
func (u *UserService) UserMenuID(userpb *pb.MenuID) (*pb.MenuItem, error) {

	result, err := menu.FetchMenuByIDHandler(u.client, userpb)
	if err != nil {
		return nil, err
	}

	return &pb.MenuItem{
		Id:        result.Id,
		Category:  result.Category,
		Name:      result.Name,
		Price:     result.Price,
		Foodimage: result.Foodimage,
		Duration:  result.Duration,
	}, nil
}

// UserMenuList implements services_inter.UserService.
// UserMenuList implements services_inter.UserService.
func (u *UserService) UserMenuList(userpb *pb.RNoparam) (*pb.MenuList, error) {
	result, err := menu.FetchByMenusHandler(u.client, userpb)
	if err != nil {
		return nil, err
	}

	var menuItems []*pb.MenuItem

	for _, i := range result.Item {
		menuItems = append(menuItems, &pb.MenuItem{
			Id:        i.Id,
			Category:  i.Category,
			Name:      i.Name,
			Price:     i.Price,
			Foodimage: i.Foodimage,
			Duration:  i.Duration,
		})
	}

	return &pb.MenuList{Item: menuItems}, nil
}
