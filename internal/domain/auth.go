package domain

type AuthPolicy struct{}

func (p *AuthPolicy) CanLogin(user *User) bool {
	return user.IsActive
}
