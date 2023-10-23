package redis

import (
	"encoding/json"
	"errors"
	"machship-go/core/entity"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisRepository(t *testing.T) {
	db, mock := redismock.NewClientMock()
	repo := NewRedisRepository(db)

	githubuser := &entity.GithubUser{Name: "kimpoytoledo"}

	data, _ := json.Marshal(githubuser)
	mock.ExpectSet("kimpoytoledo", data, CacheDuration).SetVal("OK")
	err := repo.Save("kimpoytoledo", githubuser)
	assert.NoError(t, err)

	mock.ExpectGet("kimpoytoledo").SetVal(string(data))
	result, err := repo.GetByName("kimpoytoledo")
	assert.NoError(t, err)
	assert.Equal(t, githubuser.Name, result.Name)

	mock.ExpectGet("nonexistentuser").SetErr(errors.New("not found"))
	_, err = repo.GetByName("nonexistentuser")
	assert.Error(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
