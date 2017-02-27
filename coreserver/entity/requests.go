package entity

// ReqPostUser is the body of POST /users
type ReqPostUser struct {
	// required
	Username string   `json:"username"`
	Password Password `json:"password"`
	// optional
	WeChatID string `json:"wechat_id"`
	Type     string `json:"type"`
	Email    string `json:"email"`
}

// ReqPutUser is the body of PUT /users/:id
type ReqPutUser struct {
	// optional
	Password Password `json:"password"`
	WeChatID string   `json:"wechat_id"`
	Type     string   `json:"type"`
	Email    string   `json:"email"`
}
