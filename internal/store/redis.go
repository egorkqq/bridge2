package store

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"tonconnect-bridge/internal/bridge"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr, password string, db int) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisStore{client: rdb}
}

func (r *RedisStore) Push(event *bridge.Event) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := json.Marshal(event)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal event")
		return false
	}

	// Convert event.ID to string
	eventIDStr := strconv.FormatUint(event.ID, 10)

	if err := r.client.Set(ctx, eventIDStr, data, time.Duration(event.Deadline-time.Now().Unix())*time.Second).Err(); err != nil {
		log.Error().Err(err).Msg("failed to push event to redis")
		return false
	}

	return true
}

func (r *RedisStore) ExecuteAll(lastEventId uint64, exec func(event *bridge.Event) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		data, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return err
		}

		var event bridge.Event
		if err := json.Unmarshal([]byte(data), &event); err != nil {
			return err
		}

		if event.ID <= lastEventId {
			continue
		}

		if err := exec(&event); err != nil {
			return err
		}
		log.Debug().Uint64("id", event.ID).Msg("executed event")
	}

	return nil
}

func (r *RedisStore) TestConnection(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
