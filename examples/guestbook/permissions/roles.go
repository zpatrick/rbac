package permissions

import (
	"github.com/zpatrick/rbac"
	"github.com/zpatrick/rbac/examples/guestbook/comment"
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

func NewUserRole(userID string, commentStore comment.Store, entryStore entry.Store) rbac.Role {
	return rbac.Role{
		RoleID: UserRoleID,
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("create:*", "*"),
			rbac.NewGlobPermission("list:*", "*"),
			rbac.NewGlobPermission("read:*", "*"),
			rbac.NewPermission(rbac.GlobMatch("*:comment"), ifCommentOwner(userID, commentStore)),
			rbac.NewPermission(rbac.GlobMatch("*:entry"), ifEntryOwner(userID, entryStore)),
		},
	}
}

func ifCommentOwner(userID string, store comment.Store) rbac.Matcher {
	return func(target string) (bool, error) {
		comment, err := store.GetComment(target)
		if err != nil {
			return false, nil
		}

		return comment.UserID == userID, nil
	}
}

func ifEntryOwner(userID string, store entry.Store) rbac.Matcher {
	return func(target string) (bool, error) {
		entry, err := store.GetEntry(target)
		if err != nil {
			return false, nil
		}

		return entry.UserID == userID, nil
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
