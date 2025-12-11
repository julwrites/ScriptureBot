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

func GetUser(id string, project string) User {
	ctx := context.Background()
	client := OpenClient(&ctx, project)

	var user User
	user.Id = id

	if client == nil {
		return user
	}
	defer client.Close()

	key := datastore.NameKey("User", id, nil)
	err := client.Get(ctx, key, &user)
	if err != nil {
		log.Printf("Failed to get user: %v", err)

		return user
	}

	log.Printf("Found user %s", user.Username)

	return user
}

func GetAllUsers(project string) []User {
	ctx := context.Background()
	client := OpenClient(&ctx, project)

	if client == nil {
		return []User{}
	}
	defer client.Close()

	var users []User

	_, err := client.GetAll(ctx, datastore.NewQuery("User"), &users)

	if err != nil {
		log.Printf("Failed to get users: %v", err)

		return []User{}
	}

	return users
}

func PushUser(user User, project string) bool {
	log.Printf("Updating user data %v", user)

	ctx := context.Background()
	client := OpenClient(&ctx, project)

	if client == nil {
		return false
	}
	defer client.Close()

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

func RegisterUser(platformUser def.UserData, project string) User {
	// Map identity from platform user to local user
	var user User
	user.Id = platformUser.Id
	user.Username = platformUser.Username
	user.Firstname = platformUser.Firstname
	user.Lastname = platformUser.Lastname
	user.Type = string(platformUser.Type)

	// Get stored user from DB to retrieve state (Action, Config)
	dbUser := GetUser(user.Id, project)

	// Preserve state from DB
	user.Action = dbUser.Action
	user.Config = dbUser.Config

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
