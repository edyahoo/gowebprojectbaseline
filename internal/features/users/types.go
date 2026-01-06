package users

type CreateUserRequest struct {
	Email string
	Name  string
}

type UpdateUserRequest struct {
	Name     string
	IsActive bool
}
