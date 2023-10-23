package redis

import (
	"context"
	"encoding/json"
	"machship-go/core/entity"
	"machship-go/core/port"
	"time"

	"github.com/go-redis/redis/v8"
)

const CacheDuration = 60 * time.Second

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) port.GithubRepository {
	return &redisRepository{client: client}
}

func (r *redisRepository) GetByName(name string) (*entity.GithubUser, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, name).Result()
	if err != nil {
		return nil, err
	}
	var githubuser entity.GithubUser
	err = json.Unmarshal([]byte(val), &githubuser)
	return &githubuser, err
}

func (r *redisRepository) Save(name string, githubuser *entity.GithubUser) error {
	ctx := context.Background()
	data, err := json.Marshal(githubuser)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, name, data, CacheDuration).Err()
}
