package scripturebot

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

func OpenClient(env *SessionData) *Client {
	ctx := context.Background()

	projectId := env.Secrets.PROJECT_ID

	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create Datastore client: %v", err)
		return nil
	}

	return client
}

func QueryUser(env *SessionData) UserData {
	client := OpenClient(env)

	key := datastore.NameKey("User", env.Props.User.Id, nil)

	var user UserData

	err := client.Get(ctx, key, &user)
	if err != nil {
		log.Printf("Failed to get user: %v", err)

		return user
	}

	return user
}

func UpdateUser(user *UserData, env *SessionData) bool {
	client := OpenClient(env)

	key := datastore.NameKey("User", user.Id, nil)

	_, err := client.Put(ctx, key, &user)

	if err != nil {
		log.Fatalf("Failed to put to datastore: %v", err)
		return false
	}

	return true
}
