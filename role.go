package rbac

// A Role is a grouping of permissions
type Role struct {
	RoleID      string
	Permissions []Permission
}

// Can returns true if the Role is allowed to perform the action on the target 
func (r Role) Can(action, target string) (bool, error) {
	for _, permission := range r.Permissions {
		can, err := permission(action, target)
		if err != nil {
			return false, err
		}

		if can {
			return true, nil
		}
	}

	return false, nil
}
