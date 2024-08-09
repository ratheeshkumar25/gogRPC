package di

import (
	"log"

	//"os"

	"github.com/ratheeshkumar/restaurant_user_serviceV1/config"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/db"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/handlers"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/menus"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/repositories"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/server"
	"github.com/ratheeshkumar/restaurant_user_serviceV1/internal/services"
	//"github.com/redis/go-redis/v9"
	//"github.com/twilio/twilio-go"
)

func Init() {
	config.LoadConfig()

	// Connect to the database
	dbConn := db.ConnectDB()

	// Dial the menu service client
	client, err := menus.ClientDial()
	if err != nil {
		log.Fatalf("Error dialing menu service client: %v", err)
	}

	// // Initialize Redis client
	// redisAdd := os.Getenv("REDIS_PORT")
	// redisService := config.NewRedisService(redisAdd)
	redisService, err := config.SetupRedis()
	if err != nil {
		log.Fatalf("Error initializing Redis client: %v", err)
	}

	// Initialize Twilio client
	twilioClient := config.SetupTwilio()

	// Initialize repositories and use case
	userRepo := repositories.NewUserRepo(dbConn)

	// Initialize services with all required dependencies
	userSvc := services.NewUserService(userRepo, client, redisService, twilioClient)

	// Initialize handlers and server
	userHandler := handlers.NewUserHandler(userSvc)
	server.NewGrpcServer(userHandler)
}
