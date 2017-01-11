package entity

// RunConfig is the body of POST /docker/run
type RunConfig struct {
	HostPort int    `json:"host_port"`
	Password string `json:"password"`
	// optional, in Core
	CPU float64 `json:"cpu,omitempty"`
	// optional, in MB
	Mem int `json:"mem,omitempty"`
	// optional
	Image string `json:"image,omitempty"`
	// optional
	Method string `json:"method,omitempty"`
}

// SsConfig is structure of ss client config
type SsConfig struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	Password   string `json:"password"`
	Method     string `json:"method"`
	// optional
	LocalAddress string `json:"local_address,omitempty"`
	// optional
	LocalPort int `json:"local_port,omitempty"`
	// optional
	Timeout int `json:"timeout,omitempty"`
	// optional
	FastOpen bool `json:"fast_open,omitempty"`
	// optional
	Workers int `json:"workers,omitempty"`
}
