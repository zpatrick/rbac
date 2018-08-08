package rbac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleNewGlobPermission() {
	role := Role{
		Permissions: []Permission{
			NewGlobPermission("delete:user", "*Doe"),
			NewGlobPermission("read:*", "*"),
			NewGlobPermission("*", "user_123"),
		},
	}

	fmt.Println(role.Can("read", "comment"))
	fmt.Println(role.Can("write", "books"))
	// Output:
	// [action: "delete:user"] [target: "John Doe"] => true
	// [action: "delete:user"] [target: "Jane Doe"] => true
	// [action: "delete:user"] [target: "John Smith"] => false
	// [action: "read:comment"] [target: "comment_123"] => true
	// [action: "read:article"] [target: "article_123"] => true
	// [action: "edit:user"] [target: "user_123"] => true
	// [action: "edit:user"] [target: "user_456"] => false
	// [action: "delete:user"] [target: "user_123"] => true
}

func TestGlobMatch(t *testing.T) {
	cases := map[string]map[string]bool{
		"": {
			"":        true,
			"alpha":   false,
			"beta":    false,
			"charlie": false,
		},
		"*": map[string]bool{
			"":        true,
			"alpha":   true,
			"beta":    true,
			"charlie": true,
		},
		"alpha": {
			"":        false,
			"alpha":   true,
			"beta":    false,
			"charlie": false,
		},
		"a*": {
			"":        false,
			"alpha":   true,
			"beta":    false,
			"charlie": false,
		},
		"*a": {
			"":        false,
			"alpha":   true,
			"beta":    true,
			"charlie": false,
		},
		"*a*": {
			"":        false,
			"alpha":   true,
			"beta":    true,
			"charlie": true,
		},
		"delta": {
			"":        false,
			"alpha":   false,
			"beta":    false,
			"charlie": false,
		},
	}

	for pattern, inputs := range cases {
		matcher := GlobMatch(pattern)
		for input, expected := range inputs {
			name := fmt.Sprintf("%s/%s", pattern, input)
			t.Run(name, func(t *testing.T) {
				result, err := matcher(input)
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, expected, result)
			})
		}
	}
}
