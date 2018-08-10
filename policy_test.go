package rbac

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicyRole(t *testing.T) {
	p := Policy{
		RoleID: "User",
		Permissions: map[string]string{
			"*:comment": "grid:$environment:$userID:comment:*",
		},
	}

	expected := Role{
		RoleID: "User",
		Permissions: []Permission{
			NewGlobPermission("*:comment", "grid:prod:u123:comment:*"),
		},
	}

	replacer := strings.NewReplacer("$environment", "prod", "$userID", "u123")
	assert.Equal(t, expected, p.Role(replacer))
}
