package rbac

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

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
		for _, action := range Actions {
			for _, target := range Targets {
				name := fmt.Sprintf("%s/%s:%s", c.Name, action, target)
				t.Run(name, func(t *testing.T) {
					result, err := c.Permission(action, target)
					if err != nil {
						t.Fatal(err)
					}

					c.Assert(t, action, target, result)
				})
			}
		}
	}
}
