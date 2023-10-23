package entity

type GithubUser struct {
	Name                string  `json:"name"`
	Login               string  `json:"login"`
	Company             string  `json:"company"`
	Followers           int     `json:"followers"`
	PublicRepos         int     `json:"public_repos"`
	AvgFollowersPerRepo float64 `json:"avg_followers_per_repo,omitempty"`
	Source              string  `json:"source"`
}
