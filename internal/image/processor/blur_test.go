package processor

import (
	"testing"
)

func TestBlur(t *testing.T) {
	cases := []TestCase{
		{
			Name:      "invalid blur params",
			Image:     "01.jpg",
			Params:    []string{"s"},
			ExpectErr: "invalid blur params",
		},
		{
			Name:      "invalid blur value",
			Image:     "01.jpg",
			Params:    []string{"s_100"},
			ExpectErr: "invalid blur value",
		},
		{
			Name:   "blur should be ok",
			Image:  "01.jpg",
			Params: []string{"s_2"},
		},
	}

	runTableTest(cases, t, new(Blur))
}
