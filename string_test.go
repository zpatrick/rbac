package rbac

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringMatch(t *testing.T) {
	cases := map[string]map[string]bool{
		"": {
			"":      true,
			"alpha": false,
			"beta":  false,
		},
		"alpha": {
			"":      false,
			"alpha": true,
			"beta":  false,
		},
		"beta": {
			"":      false,
			"alpha": false,
			"beta":  true,
		},
		"charlie": {
			"":      false,
			"alpha": false,
			"beta":  false,
		},
	}

	for pattern, inputs := range cases {
		matcher := StringMatch(pattern)
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
