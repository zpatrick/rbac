package rbac

// A Permission is a function that returns true if the action is allowed on the target
type Permission func(action string, target string) (bool, error)

// The Permissions type is an adapter to allow helper functions to execute on a slice of Permissions
type Permissions []Permission

// Can returns true if at least one of the permissions in p allows the action on the target
func (p Permissions) Can(action string, target string) (bool, error) {
	for _, permission := range p {
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

// NewPermission returns a Permission that will return true
// if the actionMatcher returns true for the given action, and
// if the targetMatcher returns true the given target.
func NewPermission(actionMatcher, targetMatcher Matcher) Permission {
	return func(action string, target string) (bool, error) {
		actionMatch, err := actionMatcher(action)
		if err != nil {
			return false, err
		}

		if !actionMatch {
			return false, nil
		}

		return targetMatcher(target)
	}
}

// AllowAll is a Permission that always returns true
func AllowAll(action, target string) (bool, error) {
	return true, nil
}
