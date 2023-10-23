package githubapi

import (
	"encoding/json"
	"machship-go/core/entity"
	"net/http"
)

func FetchGithubUserFromAPI(name string) (*entity.GithubUser, error) {
	resp, err := http.Get("https://api.github.com/users/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Name                string  `json:"name"`
		Login               string  `json:"login"`
		Company             string  `json:"company"`
		Followers           int     `json:"followers"`
		PublicRepos         int     `json:"public_repos"`
		AvgFollowersPerRepo float64 `json:"avg_followers_per_repo,omitempty"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &entity.GithubUser{
		Name:        result.Name,
		Login:       result.Login,
		Company:     result.Company,
		Followers:   result.Followers,
		PublicRepos: result.PublicRepos,
	}, nil
}
