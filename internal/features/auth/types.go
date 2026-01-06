package auth

type LoginRequest struct {
	Email    string
	Password string
	Remember bool
}

type LoginResponse struct {
	Success bool
	Token   string
}
