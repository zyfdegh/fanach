package entity

// Resp is the common struct of http response body
type Resp struct {
	Success bool   `json:"success"`
	ErrNo   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
}

// RespPostUser is the response to POST /user
type RespPostUser struct {
	Resp
	ID string `json:"id"`
}
