package rbac

import (
	"strings"
)

// A PolicyTemplate holds information about a Role in a templated format
type PolicyTemplate struct {
	RoleID      string            `json:"role_id"`
	Permissions map[string]string `json:"permissions"`
}

// Role converts the PolicyTemplate into a Role using GlobMatchers

/*
TODO: Investigate using custom matchers

--- namespace_admin.json ---
"role_id": "Namespace Admin",
"permissions": [
    {
        "matcher": "glob",
        "action": "*",
        "target": "grid:$namespace:*"
    }
]

--- main.go ---
var policy rbac.Policy
if err := json.Unmarshal(data, &policy); err != nil {
	return err
}

matchers := map[string]func(target, action string) rbac.Matcher {
        "glob": rbac.GlobMatch,
        "string": rbab.StringMatch,
}

replacer := strings.NewReplacer("$namespace", "development")
role, err := policy.Role(matchers, replacer)
if err != nil {
	return err
}

role.Can("delete:comment", "grid:development:u123:comment:c456")
}

*/



/*
Do something like templates?

func main() {
	p := rbac.NewPolicy("Admin").
		AddPermission("glob", "read:*", "*").
		AddPermission("glob", "*", "grid:*:$userID:*").
		// by default, has "glob", "string", and "regex", so this wouldn't need to be done every time
		SetMatcher("glob", rbac.GlobMatch).
		SetMatcher("string", rbac.StringMatch).
		SetMatcher("regex", rbac.RegexMatch)
		

	// takes (...string) for replacements
	r := p.Role("$userID", "u123")

	if err := p.WriteTo(w); err != nil {
		return err
	}

	policy, err := rbac.ReadPolicy(r)
	if err := p.WriteTo(w); err != nil {
                return err
        }

	
}

*/

func (p Policy) Role(r *strings.Replacer) Role {
	role := Role{
		RoleID:      p.RoleID,
		Permissions: make([]Permission, 0, len(p.Permissions)),
	}

	for target, action := range p.Permissions {
		permission := NewGlobPermission(r.Replace(target), r.Replace(action))
		role.Permissions = append(role.Permissions, permission)
	}

	return role
}
