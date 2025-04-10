package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/khoand3012/go-ieltsgrader/db"
)

func NewDataBase(env *Env) db.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	client, err := db.NewClient(mongodbURI)

	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	return client
}

func CloseDBConnection(client db.Client) {
	if client == nil {
		return
	}
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}
