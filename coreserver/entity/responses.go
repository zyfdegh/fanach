package entity

// Resp is the common struct of http response body
type Resp struct {
	Success bool `json:"success"`
	Err
}

// RespPostUser is the response to POST /user
type RespPostUser struct {
	Resp
	ID string `json:"id"`
}

// RespGetUser is the response to GET /user
type RespGetUser struct {
	Resp
	User User `json:"user,omitempty"`
}

// RespGetUsers is the response to GET /users
type RespGetUsers struct {
	Resp
	Users []User `json:"users,omitempty"`
}

// RespPutUser is the response to PUT /user
type RespPutUser struct {
	Resp
	User User `json:"user,omitempty"`
}
