package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringPermission(t *testing.T) {
	cases := []PermissionTestCase{
		{
			Name:       "None",
			Permission: NewStringPermission("", ""),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.False(t, result)
			},
		},
		{
			Name:       "ExactActionMatch",
			Permission: NewPermission(StringMatch("read"), Always),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, action == "read", result)
			},
		},
		{
			Name:       "ExactTargetMatch",
			Permission: NewPermission(Always, StringMatch("alpha")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, target == "alpha", result)
			},
		},
		{
			Name:       "ExactTargetsMatch",
			Permission: NewPermission(Always, StringMatch("alpha"), StringMatch("beta")),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, target == "alpha" || target == "beta", result)
			},
		},
	}

	RunPermissionsTest(t, cases...)
}
