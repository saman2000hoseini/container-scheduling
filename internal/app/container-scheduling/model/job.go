package model

type Job struct {
	Id          uint64 `json:"id"`
	Operation   string `json:"operation"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}
