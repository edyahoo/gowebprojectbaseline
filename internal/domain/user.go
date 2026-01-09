package domain

import "time"

type User struct {
	ID        UserID
	TenantID  TenantID
	Email     string
	Name      string
	Role      Role
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() error {
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if !u.TenantID.Valid() {
		return ErrInvalidTenant
	}
	return nil
}
