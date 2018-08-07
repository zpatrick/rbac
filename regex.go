package rbac

import "regexp"

// RegexMatch returns a Matcher that returns true
// if the target regular expression matches the specified pattern.
func RegexMatch(pattern string) Matcher {
	return func(target string) (bool, error) {
		return regexp.MatchString(pattern, target)
	}
}

// NewRegexPermission returns a Permission that uses RegexMatchers for the specified action and target patterns.
func NewRegexPermission(actionPattern, targetPattern string) Permission {
	return NewPermission(RegexMatch(actionPattern), RegexMatch(targetPattern))
}
