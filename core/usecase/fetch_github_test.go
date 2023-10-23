package usecase

import (
	"errors"
	"machship-go/core/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetByName(name string) (*entity.GithubUser, error) {
	args := m.Called(name)
	return args.Get(0).(*entity.GithubUser), args.Error(1)
}

func (m *MockRepository) Save(name string, githubuser *entity.GithubUser) error {
	args := m.Called(name, githubuser)
	return args.Error(0)
}

func TestFetchAndCacheGithubUser(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewGithubService(mockRepo)

	githubUser := &entity.GithubUser{Name: "Octocat"}

	// Test case: User exists in cache
	mockRepo.On("GetByName", "Octocat").Return(githubUser, nil)
	result, err := usecase.FetchAndCacheGithubUser("Octocat")
	assert.NoError(t, err)
	assert.Equal(t, githubUser, result)

	// Test case: User not in cache
	mockRepo.On("GetByName", "Octocat").Return(nil, errors.New("not found"))
	mockRepo.On("Save", "Octocat", githubUser).Return(nil)
	result, err = usecase.FetchAndCacheGithubUser("Octocat")
	assert.NoError(t, err)
	assert.Equal(t, githubUser, result)
}

func TestFetchAndCacheGithubUsers(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewGithubService(mockRepo)

	githubUser1 := &entity.GithubUser{Name: "Octocat1"}
	githubUser2 := &entity.GithubUser{Name: "Octocat2"}

	// Test case: Both users exist in cache
	mockRepo.On("GetByName", "Octocat1").Return(githubUser1, nil)
	mockRepo.On("GetByName", "Octocat2").Return(githubUser2, nil)

	results, err := usecase.FetchAndCacheGithubUsers([]string{"Octocat1", "Octocat2"})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, githubUser1)
	assert.Contains(t, results, githubUser2)

	// Test case: One user in cache, one user not in cache
	mockRepo.On("GetByName", "Octocat1").Return(nil, errors.New("not found"))
	mockRepo.On("GetByName", "Octocat2").Return(githubUser2, nil)
	mockRepo.On("Save", "Octocat1", githubUser1).Return(nil)

	results, err = usecase.FetchAndCacheGithubUsers([]string{"Octocat1", "Octocat2"})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	assert.Contains(t, results, githubUser1)
	assert.Contains(t, results, githubUser2)
}
