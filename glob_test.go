package rbac

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobPermission(t *testing.T) {
	RunPermissionsTest(t, []PermissionTestCase{
		{
			Name:       "All",
			Permission: NewGlobPermission("*", "*"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.True(t, result)
			},
		},
		{
			Name:       "None",
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
			Name:       "AnyTarget",
			Permission: NewPermission(Always, GlobMatch("*")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.True(t, result)
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
		{
			Name:       "ExactTargetsMatch",
			Permission: NewPermission(Always, GlobMatch("alpha"), GlobMatch("beta")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, target == "alpha" || target == "beta", result)
			},
		},
		{
			Name:       "GlobTargetsMatch",
			Permission: NewPermission(Always, GlobMatch("*a*"), GlobMatch("*b*")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, strings.Contains(target, "a") || strings.Contains(target, "b"), result)
			},
		},
	})
}
