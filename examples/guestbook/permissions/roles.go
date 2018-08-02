package permissions

import (
	"github.com/zpatrick/rbac"
	"github.com/zpatrick/rbac/examples/guestbook/entry"
)

const (
	AdminRoleID = "Admin"
	UserRoleID  = "User"
	GuestRoleID = "Guest"
)

func NewAdminRole() rbac.Role {
	return rbac.Role{
		RoleID: AdminRoleID,
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "*"),
		},
	}
}

// create, delete, edit, list, read
func NewUserRole(userID string, commentStore comment.Store, entryStore entry.Store) rbac.Role {
	return rbac.Role{
		RoleID: AdminRoleID,
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("create:*", "*"),
			rbac.NewGlobPermission("list:*", "*"),
			rbac.NewGlobPermission("read:*", "*"),
			rbac.NewGlobActionPermissionFunc("*:comment", ifCommentOwner(userID, commentStore)),
			rbac.NewGlobActionPermissionFunc("*:entry", ifEntryOwner(userID, entryStore)),
		},
	}
}

func ifCommentOwner(userID, store comment.Store) func(string) (bool, error) {
	return func(target string) (bool, error) {
		comment, err := store.GetComment(target)
		if err != nil {
			return false, nil
		}

	}
}

func NewGuestRole() rbac.Role {
	return rbac.Role{
		RoleID: GuestRoleID,
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("list:*", "*"),
			rbac.NewGlobPermission("read:*", "*"),
		},
	}
}
