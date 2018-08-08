package rbac

// A Matcher is a function that returns a bool representing
// whether or not the target matches some pre-defined pattern.
type Matcher func(target string) (bool, error)

// MatchAny will convert a slice of Matchers into a single Matcher
// that returns true if and only if at least one of the specified matchers returns true.
func MatchAny(matchers ...Matcher) Matcher {
	return func(target string) (bool, error) {
		for _, matcher := range matchers {
			match, err := matcher(target)
			if err != nil {
				return false, err
			}

			if match {
				return true, nil
			}
		}

		return false, nil
	}
}

// MatchAll will convert a slice of Matchers into a single Matcher
// that returns true if and only if all of the specified matchers returns true.
func MatchAll(matchers ...Matcher) Matcher {
	return func(target string) (bool, error) {
		for _, matcher := range matchers {
			match, err := matcher(target)
			if err != nil {
				return false, err
			}

			if !match {
				return false, nil
			}
		}

		return true, nil
	}
}

// Anything is a Matcher that always returns true
func Anything(target string) (bool, error) {
	return true, nil
}
