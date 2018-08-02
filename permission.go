package rbac

// A Permission is a function that returns a bool representing
// whether or not the action is allowed on the target.
type Permission func(action, target string) (bool, error)

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
