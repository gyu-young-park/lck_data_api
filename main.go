package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gyu-young-park/lck_data_api/api"
	"github.com/gyu-young-park/lck_data_api/firebaseapi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	firebaseKey := os.Getenv("FIREBASE_API_KEY")
	fmt.Println("firebaseKey:", firebaseKey)
	firebaseapi.FirebaseAppClinet = firebaseapi.NewFireBaseAPI(firebaseKey)
	server := api.NewHTTPServer()
	server.StartServer()
}
