package rbac

// A Role is a grouping of permissions.
type Role struct {
	RoleID      string
	Permissions Permissions
}

// Can returns true if the Role is allowed to perform the action on the target.
func (r Role) Can(action, target string) (bool, error) {
	return r.Permissions.Can(action, target)
}
