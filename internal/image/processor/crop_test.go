package processor

import (
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"
)

func TestCrop(t *testing.T) {
	cases := []TestCase{
		{
			Name:   "crop jpg",
			Image:  "01.jpg",
			Params: []string{"x_0", "y_0", "w_100", "h_100"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 100)
				assert.Equal(t, ref.Height(), 100)
			},
		},
		{
			Name:   "crop gif",
			Image:  "01.gif",
			Params: []string{"x_0", "y_0", "w_100", "h_100"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 100)
				assert.Equal(t, ref.PageHeight(), 100)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
	}

	runTableTest(cases, t, new(Crop))
}
