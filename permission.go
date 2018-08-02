package rbac

import "github.com/ryanuber/go-glob"

// A Permission is a function that returns a bool representing
// whether or not the action is allowed on the target.
type Permission func(action, target string) (bool, error)

// NewGlobPermission returns a Permission that will return true
// if the action glob matches the actionPattern and
// if the target glob matches one of the targetPatterns
func NewGlobPermission(actionPattern string, targetPatterns ...string) Permission {
	return func(action, target string) (bool, error) {
		if !glob.Glob(actionPattern, action) {
			return false, nil
		}

		for _, targetPattern := range targetPatterns {
			if glob.Glob(targetPattern, target) {
				return true, nil
			}
		}

		return false, nil
	}
}
