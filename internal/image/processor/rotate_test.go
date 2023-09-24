package processor

import (
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	cases := []TestCase{
		{
			Name:      "rotate with empty param",
			Image:     "01.jpg",
			Params:    []string{},
			ExpectErr: "invalid rotate params",
		},
		{
			Name:      "rotate with invalid angle",
			Image:     "01.jpg",
			Params:    []string{"500"},
			ExpectErr: "rotate angle should be in range",
		},
		{
			Name:   "rotate with 90 angle",
			Image:  "01.jpg",
			Params: []string{"90"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 267)
				assert.Equal(t, ref.Height(), 400)
			},
		},
		{
			Name:   "rotate with 45 angle",
			Image:  "01.jpg",
			Params: []string{"45"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 472)
				assert.Equal(t, ref.Height(), 472)
			},
		},
	}

	runTableTest(cases, t, new(Rotate))
}
