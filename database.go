// Brief: Database
// Primary responsibility: Key functionality needed for database

package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/google/go-cmp/cmp"
	bmul "github.com/julwrites/BotMultiplexer"
)

type UserConfig struct {
	Version       string `datastore:""`
	Timezone      string `datastore:""`
	Subscriptions string `datastore:""`
}

func OpenClient(ctx *context.Context, env *bmul.SessionData) *datastore.Client {
	projectId := env.Secrets.PROJECT_ID

	client, err := datastore.NewClient(*ctx, projectId)
	if err != nil {
		log.Printf("Failed to create Datastore client: %v", err)
		return nil
	}

	return client
}

func GetUser(env *bmul.SessionData) bmul.UserData {
	ctx := context.Background()
	client := OpenClient(&ctx, env)

	key := datastore.NameKey("User", env.User.Id, nil)

	var user bmul.UserData

	var defaultConfig UserConfig
	defaultConfig.Version = "NIV"
	UpdateUserConfig(&user, defaultConfig)

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
		log.Printf("Failed to put to datastore: %v", err)
		return false
	}

	return true
}

func GetUserConfig(user *bmul.UserData) UserConfig {
	var config UserConfig
	json.Unmarshal([]byte(user.Config), &config)

	return config
}

func UpdateUserConfig(user *bmul.UserData, config UserConfig) {
	strConfig, err := json.Marshal(config)
	if err != nil {
		log.Printf("Failed to marshal User Config: %v", err)
	}
	user.Config = string(strConfig)
}

func CompareAndUpdateUser(env *bmul.SessionData) {
	storedUser := GetUser(env)

	if !cmp.Equal(storedUser, env.User) {
		env.User.Config = storedUser.Config

		log.Printf("Updating user %s", env.User.Username)

		UpdateUser(&env.User, env)
	}
}
