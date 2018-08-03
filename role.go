package rbac

// A Role is a grouping of permissions.
type Role struct {
	RoleID      string
	Permissions Permissions
}

// Can returns true if the Role is allowed to perform the action on each of the targets.
func (r Role) Can(action string, targets ...string) (bool, error) {
	return r.Permissions.Can(action, targets...)
}
