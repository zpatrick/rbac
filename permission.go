package rbac


/*
TODO: Investigate if variadic target paramaeter is a good option.
This would allow a more dynamic set of targets to be made on an action.
For example:
	if role.Can(ListUsers)
	if role.Can(DeleteUser, userID)
	if role.Can(TransferObjectOwnership, objectID, userID)  

The associated permissions may be:
	// role can list all users
	rbac.NewGlobPermission(ListUsers)	
	
	// role can only delete current user
	rbac.NewGlobPermission(DeleteUser, userID)
	
	// role can only tranfer objects that they own, to users who are friends
	rbac.NewGlobPermission(TranferObjectOwnership, IfObjectOwner(userID), IfUserIsFriend(userID))
*/




// A Permission is a function that returns true if the action is allowed on the target
type Permission func(action string, target string) (bool, error)

// The Permissions type is an adapter to allow helper functions to execute on a slice of Permissions
type Permissions []Permission

// Can returns true if at least one of the permissions in p allows the action on the target
func (p Permissions) Can(action string, target string) (bool, error) {
	for _, permission := range p {
		can, err := permission(action, target)
		if err != nil {
			return false, err
		}

		if can {
			return true, nil
		}
	}

	return false, nil
}

// NewPermission returns a Permission that will return true
// if the actionMatcher returns true for the given action, and
// if the targetMatcher returns true the given target.
func NewPermission(actionMatcher, targetMatcher Matcher) Permission {
	return func(action string, target string) (bool, error) {
		actionMatch, err := actionMatcher(action)
		if err != nil {
			return false, err
		}

		if !actionMatch {
			return false, nil
		}

		targetMatch, err := targetMatcher(target)
		if err != nil {
			return false, err
		}

		return targetMatch, nil
	}
}
