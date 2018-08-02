package rbac

type Role struct {
	RoleID      string
	Permissions []Permission
}

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

type Roles []Role

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
