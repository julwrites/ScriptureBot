package scripturebot

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

func OpenClient(ctx *context.Context, env *SessionData) *datastore.Client {
	projectId := env.Secrets.PROJECT_ID

	client, err := datastore.NewClient(*ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create Datastore client: %v", err)
		return nil
	}

	return client
}

func QueryUser(env *SessionData) UserData {
	ctx := context.Background()
	client := OpenClient(&ctx, env)

	key := datastore.NameKey("User", env.Props.User.Id, nil)

	var user UserData

	err := client.Get(ctx, key, &user)
	if err != nil {
		log.Printf("Failed to get user: %v", err)

		return user
	}

	log.Printf("Found user %s", user.Username)

	return user
}

func UpdateUser(user *UserData, env *SessionData) bool {
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
