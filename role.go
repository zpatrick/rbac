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

// The Roles type is an adapter to allow helper functions to execute on a slice of Roles
type Roles []Role

// Can returns true if at least one of the roles in r allows the action on the target
func (r Roles) Can(action, target string) (bool, error) {
	for _, role := range r {
		can, err := role.Can(action, target)
		if err != nil {
			return false, err
		}

		if can {
			return true, nil
		}
	}

	return false, nil
}
