package rbac

// A Matcher is a function that returns a bool representing
// whether or not the target matches some pre-defined pattern.
type Matcher func(target string) (bool, error)

// Always is a Matcher that always returns true
func Always(target string) (bool, error) {
	return true, nil
}

// Never is a Matcher that always returns false
func Never(target string) (bool, error) {
	return false, nil
}
