package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Fullname   string     `json:"full_name"`
	Owner      RepoOwner  `json:"owner"`
	Permission Permission `json:"permission"`
}

type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type Permission struct {
	IsAdmin bool `json:"admin"`
	HasPush bool `json:"push"`
	HasPull bool `json:"pull"`
}
