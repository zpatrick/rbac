package rbac

// StringMatch returns a Matcher that will return a bool representing
// whether or not the target string exactly matches the pattern string.
func StringMatch(pattern string) Matcher {
	return func(target string) (bool, error) {
		return target == pattern, nil
	}
}

// NewStringPermission returns a Permission that will return true
// if the requested action exactly matches the specified action,
// and if the requested target exactly matches at least one of the specified targets.
func NewStringPermission(action string, targets ...string) Permission {
	targetMatchers := make([]Matcher, len(targets))
	for i, target := range targets {
		targetMatchers[i] = StringMatch(target)
	}

	return NewPermission(StringMatch(action), targetMatchers...)
}
