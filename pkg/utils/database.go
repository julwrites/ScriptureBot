// Brief: API for database handling
// Primary responsibility: API layer between GCloud datastore and other functionality

package utils

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/julwrites/BotPlatform/pkg/def"
)

type UserConfig struct {
	Version       string
	Timezone      string
	Subscriptions string
}

func OpenClient(ctx *context.Context, project string) *firestore.Client {
	client, err := firestore.NewClient(*ctx, project)
	if err != nil {
		log.Printf("Failed to create Firestore client: %v", err)
		return nil
	}

	return client
}

func GetUser(user def.UserData, project string) def.UserData {
	ctx := context.Background()
	client := OpenClient(&ctx, project)

	doc, err := client.Collection("User").Doc(user.Id).Get(ctx)
	if err != nil {
		log.Printf("Failed to get user doc: %v", err)

		return user
	}

	var obj def.UserData
	err = doc.DataTo(&obj)

	if err != nil {
		log.Printf("Failed to unmarshal user data: %v", err)

		return user
	}

	user = obj

	log.Printf("Found user %s", user.Username)

	return user
}

func PushUser(user def.UserData, project string) bool {
	log.Printf("Updating user data %v", user)

	ctx := context.Background()
	client := OpenClient(&ctx, project)

	_, err := client.Collection("User").Doc(user.Id).Set(ctx, user)

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

func RegisterUser(user def.UserData, project string) def.UserData {
	// Get stored user if any, else default to what we currently have
	user = GetUser(user, project)

	// Read the stored config
	config := DeserializeUserConfig(user.Config)
	// If storedconfig is not complete, set the default data
	if len(config.Version) == 0 {
		config.Version = "NIV"
	}

	user.Config = SerializeUserConfig(config)

	log.Printf("User's current state: %v", user)

	return user
}
