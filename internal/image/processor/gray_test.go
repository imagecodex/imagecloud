package processor

import (
	"testing"
)

func TestGray(t *testing.T) {
	cases := []TestCase{
		{
			Name:      "invalid params",
			Image:     "01.jpg",
			Params:    []string{},
			ExpectErr: "invalid gray params",
		},
		{
			Name:   "enable gray",
			Image:  "01.jpg",
			Params: []string{"1"},
		},
		{
			Name:   "disable gray",
			Image:  "01.jpg",
			Params: []string{"0"},
		},
	}

	runTableTest(cases, t, new(Gray))
}
