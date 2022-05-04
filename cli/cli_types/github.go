package cli_types

type GhFileTreeResponse struct {
	SHA       string   `json:"sha"`
	URL       string   `json:"url"`
	Tree      []GhTree `json:"tree"`
	Truncated bool     `json:"truncated"`
}

type GhTree struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	SHA  string `json:"sha"`
	Size *int64 `json:"size,omitempty"`
	URL  string `json:"url"`
}

type GhIssue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Labels []string `json:"labels"`
}

type GhIssueResponse struct {
	Id  int    `json:"number"`
	Url string `json:"html_url"`
}

type GhFile struct {
	SHA      string `json:"sha"`
	NodeID   string `json:"node_id"`
	Size     int64  `json:"size"`
	URL      string `json:"url"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}
