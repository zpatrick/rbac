package rbac

// StringMatch returns a Matcher that returns true
// if the target is equal to the specified pattern.
func StringMatch(pattern string) Matcher {
	return func(target string) (bool, error) {
		return target == pattern, nil
	}
}

// NewStringPermission returns a Permission that uses StringMatchers for the specified action and targets.
func NewStringPermission(action, target string) Permission {
	return NewPermission(StringMatch(action), StringMatch(target))
}
