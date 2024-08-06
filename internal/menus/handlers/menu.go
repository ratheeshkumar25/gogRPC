package handlers

import (
	"context"
	"log"

	menupb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/menus/pb"
	pb "github.com/ratheeshkumar/restaurant_user_serviceV1/internal/pb"
)

func FetchMenuByIDHandler(client menupb.MenuServiceClient, p *pb.MenuID) (*menupb.MenuItem, error) {
	ctx := context.Background()
	response, err := client.FetchMenuByID(ctx, &menupb.MenuID{Id: p.Id})
	if err != nil {
		log.Printf("error while fetching the menu by ID")
		return nil, err
	}
	return response, nil
}

func FetchMenuBYNameHandler(client menupb.MenuServiceClient, p *pb.FoodByName) (*menupb.MenuItem, error) {
	ctx := context.Background()
	response, err := client.FetchMenuByName(ctx, &menupb.FoodByName{Name: p.Name})
	if err != nil {
		log.Printf("error while fetching the menu by Name")
		return nil, err
	}
	return response, nil
}

func FetchByMenusHandler(client menupb.MenuServiceClient, p *pb.RNoparam) (*menupb.MenuList, error) {
	ctx := context.Background()
	response, err := client.FetchMenus(ctx, &menupb.NoParam{})
	if err != nil {
		log.Printf("errror while fetching the menulist")
		return nil, err
	}
	return response, nil
}