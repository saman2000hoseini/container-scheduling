package request

type JobRequest struct {
	Operation   string `json:"operation"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}
