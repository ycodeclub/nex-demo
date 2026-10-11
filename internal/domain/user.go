package domain

import (
	"errors"
	"strings"
)

// User 是演示用用户聚合。
type User struct {
	id          string
	username    string
	password    string // 明文仅用于 demo；生产请存哈希
	displayName string
	roles       []string
}

// NewUser 创建用户。
func NewUser(id, username, password, displayName string, roles []string) (*User, error) {
	id = strings.TrimSpace(id)
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	displayName = strings.TrimSpace(displayName)
	if id == "" || username == "" || password == "" {
		return nil, errors.New("用户 id/用户名/密码不能为空")
	}
	if displayName == "" {
		displayName = username
	}
	if len(roles) == 0 {
		roles = []string{"user"}
	}
	return &User{
		id:          id,
		username:    username,
		password:    password,
		displayName: displayName,
		roles:       append([]string(nil), roles...),
	}, nil
}

// ID 返回用户标识。
func (u *User) ID() string {
	if u == nil {
		return ""
	}
	return u.id
}

// Username 返回登录名。
func (u *User) Username() string {
	if u == nil {
		return ""
	}
	return u.username
}

// DisplayName 返回展示名。
func (u *User) DisplayName() string {
	if u == nil {
		return ""
	}
	return u.displayName
}

// Roles 返回角色列表副本。
func (u *User) Roles() []string {
	if u == nil {
		return nil
	}
	return append([]string(nil), u.roles...)
}

// MatchPassword 校验密码（demo 明文比对）。
func (u *User) MatchPassword(password string) bool {
	return u != nil && u.password == password
}
