package auth_structs

type LoginRequest struct {
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
