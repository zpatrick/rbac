package rbac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RoleTestCase struct {
	Action   string
	Target   string
	Expected bool
}

func RunRoleTest(t *testing.T, role Role, cases ...RoleTestCase) {
	for _, c := range cases {
		name := fmt.Sprintf("%s/%s(%s)", role.RoleID, c.Action, c.Target)
		t.Run(name, func(t *testing.T) {
			result, err := role.Can(c.Action, c.Target)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, result, c.Expected)
		})
	}
}
