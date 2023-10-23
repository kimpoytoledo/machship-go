package usecase

import (
	"machship-go/adapter/githubapi"
	"machship-go/core/entity"
	"machship-go/core/port"
	"sort"
)

type githubService struct {
	repo port.GithubRepository
}

func NewGithubService(r port.GithubRepository) port.GithubUsecase {
	return &githubService{repo: r}
}

func (s *githubService) FetchAndCacheGithubUser(name string) (*entity.GithubUser, error) {
	githubUser, err := s.repo.GetByName(name)
	if err == nil {
		githubUser.Source = "Redis Cache"
		return githubUser, nil
	}

	githubUser, err = githubapi.FetchGithubUserFromAPI(name)
	if err != nil || githubUser.Name == "" {
		return &entity.GithubUser{
			Name:   name,
			Source: "User not found in Github",
		}, nil
	}
	githubUser.Source = "githubAPI"

	if githubUser.PublicRepos > 0 {
		githubUser.AvgFollowersPerRepo = float64(githubUser.Followers / githubUser.PublicRepos)
	} else {
		githubUser.AvgFollowersPerRepo = 0
	}

	if err := s.repo.Save(name, githubUser); err != nil {
		return nil, err
	}

	return githubUser, nil
}

func (s *githubService) FetchAndCacheGithubUsers(names []string) ([]*entity.GithubUser, error) {
	var results []*entity.GithubUser

	for _, name := range names {
		user, err := s.FetchAndCacheGithubUser(name)
		if err != nil {
			user = &entity.GithubUser{
				Name:   name,
				Source: "Error: " + err.Error(),
			}
		}
		results = append(results, user)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})

	return results, nil
}
