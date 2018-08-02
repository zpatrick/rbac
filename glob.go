package rbac

import "github.com/ryanuber/go-glob"

// GlobMatch returns a Matcher that returns true
// if the target glob matches the specified pattern.
func GlobMatch(pattern string) Matcher {
	return func(target string) (bool, error) {
		return glob.Glob(pattern, target), nil
	}
}

// NewGlobPermission returns a Permission that uses GlobMatchers for the specified action and target patterns.
func NewGlobPermission(actionPattern string, targetPatterns ...string) Permission {
	targetMatchers := make([]Matcher, len(targetPatterns))
	for i, targetPattern := range targetPatterns {
		targetMatchers[i] = GlobMatch(targetPattern)
	}

	return NewPermission(GlobMatch(actionPattern), targetMatchers...)
}
