// Brief: Database
// Primary responsibility: Key functionality needed for database

package main

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/google/go-cmp/cmp"
	bmul "github.com/julwrites/BotMultiplexer"
)

func OpenClient(ctx *context.Context, env *bmul.SessionData) *datastore.Client {
	projectId := env.Secrets.PROJECT_ID

	client, err := datastore.NewClient(*ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create Datastore client: %v", err)
		return nil
	}

	return client
}

func QueryUser(env *bmul.SessionData) bmul.UserData {
	ctx := context.Background()
	client := OpenClient(&ctx, env)

	key := datastore.NameKey("User", env.User.Id, nil)

	var user bmul.UserData

	err := client.Get(ctx, key, &user)
	if err != nil {
		log.Printf("Failed to get user: %v", err)

		return user
	}

	log.Printf("Found user %s", user.Username)

	return user
}

func UpdateUser(user *bmul.UserData, env *bmul.SessionData) bool {
	ctx := context.Background()
	client := OpenClient(&ctx, env)

	key := datastore.NameKey("User", user.Id, nil)

	_, err := client.Put(ctx, key, user)

	if err != nil {
		log.Fatalf("Failed to put to datastore: %v", err)
		return false
	}

	return true
}

func CompareAndUpdateUser(env *bmul.SessionData) {
	storedUser := QueryUser(env)

	if !cmp.Equal(storedUser, env.User) {
		env.User.Config = storedUser.Config

		log.Printf("Updating user %s", env.User.Username)

		UpdateUser(&env.User, env)
	}
}
