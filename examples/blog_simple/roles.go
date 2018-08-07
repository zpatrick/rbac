package main

import "github.com/zpatrick/rbac"

// NewAdminRole returns a role with admin-level permissions
func NewAdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "*"),
		},
	}
}

// NewGuestRole returns a role with guest-level permissions
func NewGuestRole() rbac.Role {
	return rbac.Role{
		RoleID: "Guest",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("ReadArticle", "*"),
			rbac.NewGlobPermission("RateArticle", "*"),
		},
	}
}
