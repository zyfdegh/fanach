package entity

// ReqPostRm is the body of POST /docker/rm request
type ReqPostRm struct {
	// ID or name of container
	ID string `json:"id"`
}

// ReqPostRun is the body of POST /docker/run request
type ReqPostRun struct {
	RunConfig
}
