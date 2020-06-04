// Brief: API for database handling
// Primary responsibility: API layer between GCloud datastore and other functionality

package api

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/datastore"
	bmul "github.com/julwrites/BotMultiplexer"
)

type UserConfig struct {
	Version       string `datastore:""`
	Timezone      string `datastore:""`
	Subscriptions string `datastore:""`
}

func OpenClient(ctx *context.Context, env bmul.SessionData) *datastore.Client {
	projectId := env.Secrets.PROJECT_ID

	client, err := datastore.NewClient(*ctx, projectId)
	if err != nil {
		log.Printf("Failed to create Datastore client: %v", err)
		return nil
	}

	return client
}

func GetUser(env bmul.SessionData) bmul.UserData {
	ctx := context.Background()
	client := OpenClient(&ctx, env)

	key := datastore.NameKey("User", env.User.Id, nil)

	var user bmul.UserData

	err := client.Get(ctx, key, &user)
	if err != nil {
		log.Printf("Failed to get user: %v", err)

		return env.User
	}

	log.Printf("Found user %s", user.Username)

	return user
}

func PushUser(env bmul.SessionData) bool {
	log.Printf("Updating user data %v", env.User)

	ctx := context.Background()
	client := OpenClient(&ctx, env)

	key := datastore.NameKey("User", env.User.Id, nil)

	_, err := client.Put(ctx, key, &env.User)

	if err != nil {
		log.Printf("Failed to put to datastore: %v", err)
		return false
	}

	return true
}

func DeserializeUserConfig(config string) UserConfig {
	var userConfig UserConfig
	err := json.Unmarshal([]byte(config), &userConfig)
	if err != nil {
		log.Printf("Failed to unmarshal User Config: %v", err)
	}
	return userConfig
}

func SerializeUserConfig(config UserConfig) string {
	strConfig, err := json.Marshal(config)
	if err != nil {
		log.Printf("Failed to marshal User Config: %v", err)
	}

	return string(strConfig)
}

func RegisterUser(env bmul.SessionData) bmul.SessionData {
	// Get stored user if any, else default to what we currently have
	env.User = GetUser(env)

	// Read the stored config
	config := DeserializeUserConfig(env.User.Config)
	// If stored config is not complete, set the default data
	if len(config.Version) == 0 {
		config.Version = "NIV"
	}

	env.User.Config = SerializeUserConfig(config)

	log.Printf("User's current state: %v", env.User)

	return env
}
