// Brief: API for database handling
// Primary responsibility: API layer between GCloud datastore and other functionality

package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/julwrites/BotPlatform/pkg/def"
)

type UserConfig struct {
	Version       string
	Timezone      string
	Subscriptions string
}

func write(client *storage.Client, bucket, object string, data []byte) error {
	// [START upload_file]
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err := wc.Write(data); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	// [END upload_file]
	return nil
}

func read(client *storage.Client, bucket, object string) ([]byte, error) {
	// [START download_file]
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
	// [END download_file]
}

func OpenClient(ctx *context.Context) *storage.Client {
	client, err := storage.NewClient(*ctx)
	if err != nil {
		log.Printf("Failed to create Datastore client: %v", err)
		return nil
	}

	return client
}

func GetUser(user def.UserData, bucket string) def.UserData {
	ctx := context.Background()
	client := OpenClient(&ctx)
	key := bucket
	object := user.Id

	data, err := read(client, key, object)
	if err != nil {
		log.Printf("Failed to get user data: %v", err)

		return user
	}

	var obj def.UserData
	err = json.Unmarshal(data, &obj)

	if err != nil {
		log.Printf("Failed to unmarshal user data: %v", err)

		return user
	}

	user = obj

	log.Printf("Found user %s", user.Username)

	return user
}

func PushUser(user def.UserData, bucket string) bool {
	log.Printf("Updating user data %v", user)

	ctx := context.Background()
	client := OpenClient(&ctx)
	key := bucket
	object := user.Id

	data, err := json.Marshal(user)
	if err != nil {
		log.Printf("Failed to marshal User Config: %v", err)
		return false
	}

	err = write(client, key, object, []byte(data))

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

func RegisterUser(user def.UserData, bucket string) def.UserData {
	// Get stored user if any, else default to what we currently have
	user = GetUser(user, bucket)

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
