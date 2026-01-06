package domain

type UserID int64

type TenantID int64

func (id UserID) Valid() bool {
	return id > 0
}

func (id TenantID) Valid() bool {
	return id > 0
}
