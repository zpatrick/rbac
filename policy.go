package rbac

import (
	"encoding/json"
	"fmt"
	"strings"
)

// A PolicyTemplate holds information about a Role in a templated format
type PolicyTemplate struct {
	RoleID              string                                            `json:"role_id"`
	PermissionTemplates []PermissionTemplate                              `json:"permissions"`
	constructors        map[string]func(action, target string) Permission `json:"-"`
}

// A PermissionTemplate holds information about a permission in templated format.
type PermissionTemplate struct {
	Constructor string `json:"constructor"`
	Action      string `json:"action"`
	Target      string `json:"target"`
}

// NewPolicyTemplate generates a new PolicyTemplate with the specified roleID.
func NewPolicyTemplate(roleID string) *PolicyTemplate {
	return &PolicyTemplate{
		RoleID:              roleID,
		PermissionTemplates: []PermissionTemplate{},
		constructors:        DefaultConstructors(),
	}
}

func (p *PolicyTemplate) AddPermission(constructor, action, target string) *PolicyTemplate {
	p.PermissionTemplates = append(p.PermissionTemplates, PermissionTemplate{constructor, action, target})
	return p
}

func (p *PolicyTemplate) SetConstructor(name string, constructor func(action, target string) Permission) *PolicyTemplate {
	p.constructors[name] = constructor
	return p
}

func DefaultConstructors() map[string]func(action, target string) Permission {
	return map[string]func(action, target string) Permission{
		"glob":   NewGlobPermission,
		"regex":  NewRegexPermission,
		"string": NewStringPermission,
	}
}

func (p *PolicyTemplate) Role(oldnew ...string) (*Role, error) {
	role := &Role{
		RoleID:      p.RoleID,
		Permissions: make(Permissions, len(p.PermissionTemplates)),
	}

	replacer := strings.NewReplacer(oldnew...)
	for i, permissionTemplate := range p.PermissionTemplates {
		constructor, ok := p.constructors[permissionTemplate.Constructor]
		if !ok {
			return nil, fmt.Errorf("No constructor set for '%s'", permissionTemplate.Constructor)
		}

		action := replacer.Replace(permissionTemplate.Action)
		target := replacer.Replace(permissionTemplate.Target)
		role.Permissions[i] = constructor(action, target)
	}

	return role, nil
}

func (p *PolicyTemplate) UnmarshalJSON(data []byte) error {
	type Alias PolicyTemplate
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.constructors = DefaultConstructors()
	return nil
}
