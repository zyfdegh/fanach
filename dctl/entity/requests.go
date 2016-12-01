package entity

// ReqPostRm is the body of POST /docker/rm request
type ReqPostRm struct {
	// ID or name of container
	ID string `json:"id"`
}

// ReqPostRun is the body of POST /docker/run request
type ReqPostRun struct {
	HostPort string `json:"hostPort"`
	Password string `json:"password"`
	// optional
	CPU float64 `json:"cpu"`
	// optional, in MB
	Mem int `json:"mem"`
	// optional
	Image string `json:"image"`
	// optional
	Method string `json:"method"`
}
