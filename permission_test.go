package rbac

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Actions = []string{
	"create",
	"read",
	"update",
	"delete",
}

var Targets = []string{
	"alpha",
	"bravo",
	"charlie",
	"delta",
	"echo",
	"foxtrot",
	"golf",
	"hotel",
	"india",
	"juliett",
	"kilo",
	"lima",
	"mike",
	"november",
	"oscar",
	"papa",
	"quebec",
	"romeo",
	"sierra",
	"tango",
	"uniform",
	"victor",
	"whiskey",
	"x-ray",
	"yankee",
	"zulu",
}

type PermissionTestCase struct {
	Name       string
	Permission Permission
	Assert     func(t *testing.T, action, target string, result bool)
}

func RunPermissionsTest(t *testing.T, cases []PermissionTestCase) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			for _, action := range Actions {
				for _, target := range Targets {
					t.Run(fmt.Sprintf("%s:%s", action, target), func(t *testing.T) {
						result, err := c.Permission(action, target)
						if err != nil {
							t.Fatal(err)
						}

						c.Assert(t, action, target, result)
					})
				}
			}
		})
	}
}

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
			Name:       "ExactActionMatch",
			Permission: NewGlobPermission("read", "*"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, action == "read", result)
			},
		},
		{
			Name:       "GlobActionMatch",
			Permission: NewGlobPermission("*a*", "*"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, strings.Contains(action, "a"), result)
			},
		},
		{
			Name:       "ExactTargetMatch",
			Permission: NewGlobPermission("*", "alpha"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, target == "alpha", result)
			},
		},
		{
			Name:       "GlobTargetMatch",
			Permission: NewGlobPermission("*", "*a*"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, strings.Contains(target, "a"), result)
			},
		},
		{
			Name:       "ExactTargetsMatch",
			Permission: NewGlobPermission("*", "alpha", "beta", "charlie"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, target == "alpha" ||
					target == "beta" ||
					target == "charlie",
					result)
			},
		},
		{
			Name:       "GlobTargetsMatch",
			Permission: NewGlobPermission("*", "*a*", "*b*", "*c*"),
			Assert: func(t *testing.T, action, target string, result bool) {
				assert.Equal(t, strings.Contains(target, "a") ||
					strings.Contains(target, "b") ||
					strings.Contains(target, "c"),
					result)
			},
		},
	})
}
