package request

type Job_request struct {
	Id          uint64
	Operation   string
	Source      string
	Destination string
}
