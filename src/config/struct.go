package config

// for json storage
type Server struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

type Account struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

// session
type Session struct {
	Name string `json:"name,omitempty"`
}
