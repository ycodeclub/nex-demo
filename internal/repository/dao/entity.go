package dao

// UserRow 是持久化行模型（demo 仅作分层占位）。
type UserRow struct {
	ID          string
	Username    string
	Password    string
	DisplayName string
	Roles       []string
}
