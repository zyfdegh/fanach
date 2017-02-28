package entity

// Resp is the common struct of http response body
type Resp struct {
	Success bool `json:"success"`
	Err
}

// RespPostUser is the response to POST /users
type RespPostUser struct {
	Resp
	User User `json:"user,omitempty"`
}

// RespGetUser is the response to GET /users/:id
type RespGetUser struct {
	Resp
	User User `json:"user,omitempty"`
}

// RespGetUsers is the response to GET /users
type RespGetUsers struct {
	Resp
	Users []User `json:"users,omitempty"`
}

// RespPutUser is the response to PUT /users/:id
type RespPutUser struct {
	Resp
	User User `json:"user,omitempty"`
}

// RespPostSess is the response to POST /sess
type RespPostSess struct {
	Resp
	Sess Session `json:"sess,omitempty"`
}
