package rbac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
