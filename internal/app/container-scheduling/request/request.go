package request

type JobRequest struct {
	Id          uint64
	Operation   string
	Source      string
	Destination string
}
