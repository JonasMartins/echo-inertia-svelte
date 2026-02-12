// Package users ...
package users

import (
	"echo-inertia.com/src/internal/domain/role"
	base "echo-inertia.com/src/pkg/models"
)

type User struct {
	Base     base.Base
	Name     string
	Email    string
	Password string
	Role     *role.Role
	RoleType string
	Enabled  bool
}
