package gopgtesting

import (
	"context"
	"testing"
)

func TestDBConnection(t *testing.T) {
	db := NewDBConnection()
	ctx := context.Background()

	if err := db.Ping(ctx); err != nil {
		t.Error(err)
	}
}
