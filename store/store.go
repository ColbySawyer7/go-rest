package store

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// IEventStore is the database interface for storage of Event objects
type IEventStore interface {
	Get(ctx context.Context, in *object.Event) (*object.Event, error)
	List(ctx context.Context, in *object.Event) (*object.Event, error)
	Create(ctx context.Context, in *object.Event) (*object.Event, error)
	UpdateDetails(ctx context.Context, in *object.Event) (*object.Event, error)
	Cancel(ctx context.Context, in *object.Event) (*object.Event, error)
	Reschedule(ctx context.Context, in *object.Event) (*object.Event, error)
	Delete(ctx context.Context, in *object.Event) (*object.Event, error)
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

// GenerateUniqueID returns a time based sortable unique identifier
func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})
	now := time.Now().UTC()
	return fmt.Sprintf("%010v-%010v-%s", now.Unix(), now.Nanosecond(), string(word))
}
