package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot load enc variable")
		panic(err)
	}
	databaseUrl := os.Getenv("DATABASE_URL")
	client, err := mongo.Connect(context.TODO(), options.Client(), options.Client().ApplyURI(databaseUrl))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", client)
	fmt.Println("this is from database package")
}
