package rbac

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobPermission(t *testing.T) {
	cases := []PermissionTestCase{
		{
			Name:       "AllowAll",
			Permission: NewGlobPermission("*", "*"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.True(t, result)
			},
		},
		{
			Name:       "AllowNone",
			Permission: NewGlobPermission("", ""),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.False(t, result)
			},
		},
		{
			Name:       "AnyAction",
			Permission: NewPermission(GlobMatch("*"), Always),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.True(t, result)
			},
		},
		{
			Name:       "NoAction",
			Permission: NewPermission(GlobMatch(""), Always),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.False(t, result)
			},
		},
		{
			Name:       "AnyTarget",
			Permission: NewPermission(Always, GlobMatch("*")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.True(t, result)
			},
		},
		{
			Name:       "NoTarget",
			Permission: NewPermission(Always, GlobMatch("")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.False(t, result)
			},
		},
		{
			Name:       "ExactActionMatch",
			Permission: NewPermission(GlobMatch("read"), Always),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, action == "read", result)
			},
		},
		{
			Name:       "GlobActionMatch",
			Permission: NewPermission(GlobMatch("*a*"), Always),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, strings.Contains(action, "a"), result)
			},
		},
		{
			Name:       "ExactTargetMatch",
			Permission: NewPermission(Always, GlobMatch("alpha")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, target == "alpha", result)
			},
		},
		{
			Name:       "GlobTargetMatch",
			Permission: NewPermission(Always, GlobMatch("*a*")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, strings.Contains(target, "a"), result)
			},
		},
	}

	RunPermissionsTest(t, cases...)
}
