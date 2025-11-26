// Brief: API for database handling
// Primary responsibility: API layer between GCloud datastore and other functionality

package utils

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"cloud.google.com/go/datastore"
	"github.com/julwrites/BotPlatform/pkg/def"
	"github.com/julwrites/ScriptureBot/pkg/secrets"
)

var (
	cachedDatabaseID     string
	cachedDatabaseIDOnce sync.Once
)

func GetDatabaseID() string {
	cachedDatabaseIDOnce.Do(func() {
		id, err := secrets.Get("USER_DATABASE_ID")
		if err == nil && id != "" {
			cachedDatabaseID = id
			return
		}

		log.Printf("Warning: USER_DATABASE_ID not found, defaulting to 'scripturebot-users'. Error: %v", err)
		cachedDatabaseID = "scripturebot-users"
	})
	return cachedDatabaseID
}

type UserConfig struct {
	Version       string `datastore:""`
	Timezone      string `datastore:""`
	Subscriptions string `datastore:""`
}

func OpenClient(ctx *context.Context, project string) *datastore.Client {
	dbID := GetDatabaseID()
	client, err := datastore.NewClientWithDatabase(*ctx, project, dbID)
	if err != nil {
		log.Printf("Failed to create Firestore client: %v", err)
		return nil
	}

	return client
}

func GetUser(user def.UserData, project string) def.UserData {
	ctx := context.Background()
	client := OpenClient(&ctx, project)

	key := datastore.NameKey("User", user.Id, nil)
	var entity def.UserData
	err := client.Get(ctx, key, &entity)
	if err != nil {
		log.Printf("Failed to get user: %v", err)

		return user
	}

	user = entity

	log.Printf("Found user %s", user.Username)

	return user
}

func GetAllUsers(project string) []def.UserData {
	ctx := context.Background()
	client := OpenClient(&ctx, project)

	var users []def.UserData

	_, err := client.GetAll(ctx, datastore.NewQuery("User"), &users)

	if err != nil {
		log.Printf("Failed to get users: %v", err)

		return []def.UserData{}
	}

	return users
}

func PushUser(user def.UserData, project string) bool {
	log.Printf("Updating user data %v", user)

	ctx := context.Background()
	client := OpenClient(&ctx, project)

	key := datastore.NameKey("User", user.Id, nil)

	_, err := client.Put(ctx, key, &user)

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
