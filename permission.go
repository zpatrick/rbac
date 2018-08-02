package rbac

// A Permission is a function that returns true if the action is allowed on the target
type Permission func(action, target string) (bool, error)

// NewPermission returns a Permission that will return true
// if the actionMatcher returns true for the given action,
// and at least one of the targetMatchers returns true for the given target.
func NewPermission(actionMatcher Matcher, targetMatchers ...Matcher) Permission {
	return func(action, target string) (bool, error) {
		actionMatch, err := actionMatcher(action)
		if err != nil {
			return false, err
		}

		if !actionMatch {
			return false, nil
		}

		for _, targetMatcher := range targetMatchers {
			targetMatch, err := targetMatcher(target)
			if err != nil {
				return false, err
			}

			if targetMatch {
				return true, nil
			}
		}

		return false, nil
	}
}
