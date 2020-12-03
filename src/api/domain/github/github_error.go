package github

type GithubErrorResponse struct {
	StatusCode       int
	Message          string        `json:"message"`
	DocumentationUrl string        `json:"documentation_url"`
	Error            []GithubError `json:"error"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
