package entity

// Resp is the common struct of http response body
type Resp struct {
	Success bool   `json:"success"`
	ErrNo   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
}

// RespPostRm is the response to POST /docker/rm
type RespPostRm struct {
	Resp
}

// RespPostRun is the response to POST /docker/run
type RespPostRun struct {
	Resp
	ID string `json:"id"`
}

// RespGetStats is the response to GET /docker/stats
type RespGetStats struct {
	Resp
	Stats
}
