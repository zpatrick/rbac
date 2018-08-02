package rbac

import "github.com/ryanuber/go-glob"

// GlobMatch returns a Matcher that will return a bool representing
// whether or not the target glob matches the pattern.
func GlobMatch(pattern string) Matcher {
	return func(target string) (bool, error) {
		return glob.Glob(pattern, target), nil
	}
}

// NewGlobPermission returns a Permission that will return true
// if the requested action glob matches the specified actionPattern,
// and if the requested target glob matches one of the specified targetPatterns.
func NewGlobPermission(actionPattern string, targetPatterns ...string) Permission {
	targetMatchers := make([]Matcher, len(targetPatterns))
	for i, targetPattern := range targetPatterns {
		targetMatchers[i] = GlobMatch(targetPattern)
	}

	return NewPermission(GlobMatch(actionPattern), targetMatchers...)
}
