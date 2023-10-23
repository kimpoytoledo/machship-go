package port

import "machship-go/core/entity"

type GithubRepository interface {
	GetByName(name string) (*entity.GithubUser, error)
	Save(name string, githubuser *entity.GithubUser) error
}

type GithubUsecase interface {
	FetchAndCacheGithubUser(name string) (*entity.GithubUser, error)
	FetchAndCacheGithubUsers(names []string) ([]*entity.GithubUser, error)
}
