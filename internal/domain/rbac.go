package domain

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleUser   Role = "user"
	RoleViewer Role = "viewer"
)

type Permission string

const (
	PermissionReadUsers  Permission = "users:read"
	PermissionWriteUsers Permission = "users:write"
)
