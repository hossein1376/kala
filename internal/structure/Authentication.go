package structure

type LoginRequest struct {
	Username string `json:"username"  example:"username"`
	Password string `json:"password"  example:"password"`
}

type LoginResponse struct {
	Token string `json:"token"  example:"random_jwt_token"`
}
