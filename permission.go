package rbac

// A Permission is a function that returns true if the action is allowed on all of the targets
type Permission func(action string, targets ...string) (bool, error)

// The Permissions type is an adapter to allow helper functions to execute on a slice of Permissions
type Permissions []Permission

// Can returns true if at least one of the permissions in p allows the action on all of the targets
func (p Permissions) Can(action string, targets ...string) (bool, error) {
	for _, permission := range p {
		can, err := permission(action, targets...)
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
// if the actionMatcher returns true for the given action,
// and at least one of the targetMatchers returns true for each of the given targets.
func NewPermission(actionMatcher Matcher, targetMatchers ...Matcher) Permission {
	return func(action string, targets ...string) (bool, error) {
		actionMatch, err := actionMatcher(action)
		if err != nil {
			return false, err
		}

		if !actionMatch {
			return false, nil
		}

		for _, target := range targets {
			var targetHasMatch bool
			for _, targetMatcher := range targetMatchers {
				targetMatch, err := targetMatcher(target)
				if err != nil {
					return false, err
				}

				if targetMatch {
					targetHasMatch = true
					break
				}
			}

			if !targetHasMatch {
				return false, nil
			}
		}

		return true, nil
	}
}
