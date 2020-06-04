package api

import (
	"context"
	"testing"

	bmul "github.com/julwrites/BotMultiplexer"
)

func TestOpenClient(t *testing.T) {
	var data bmul.SessionData

	ctx := context.Background()
	client := OpenClient(&ctx, data)

	if client == nil {
		t.Errorf("Could not open client")
	}
}
