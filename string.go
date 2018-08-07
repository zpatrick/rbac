package rbac

// StringMatch returns a Matcher that returns true
// if the target string matches the specified pattern.
func StringMatch(pattern string) Matcher {
	return func(target string) (bool, error) {
		return pattern == target, nil
	}
}

// NewStringPermission returns a Permission that uses StringMatchers for the specified action and target.
func NewStringPermission(action, target string) Permission {
	return NewPermission(StringMatch(action), StringMatch(target))
}
