package processor

import (
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"
)

func TestAutoRotate(t *testing.T) {
	ref, err := loadImage("f.jpeg")
	assert.Nil(t, err)
	oInfo := ref.Metadata()
	ref.Close()

	cases := []TestCase{
		{
			Name:   "auto orient disable",
			Image:  "f.jpeg",
			Params: []string{"0"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), oInfo.Width)
				assert.Equal(t, ref.Height(), oInfo.Height)
			},
		},
		{
			Name:   "auto orient enable",
			Image:  "f.jpeg",
			Params: []string{"1"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), oInfo.Height)
				assert.Equal(t, ref.Height(), oInfo.Width)
			},
		},
	}

	runTableTest(cases, t, new(AutoRotate))
}
