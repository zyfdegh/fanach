package main

// SsAcount is structure of shadowsocks config file
type SsAcount struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	Password   string `json:"password"`
	Method     string `json:"method"`
}
