package processor

import (
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"
)

func TestResize(t *testing.T) {
	cases := []TestCase{
		{
			Name:   "resize with lfit",
			Image:  "01.jpg",
			Params: []string{"m_lfit", "w_100", "h_200"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 100)
				assert.Equal(t, ref.Height(), 66)
			},
		},
		{
			Name:   "resize with mfit",
			Image:  "01.jpg",
			Params: []string{"m_mfit", "w_100", "h_200"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 299)
				assert.Equal(t, ref.Height(), 200)
			},
		},
		{
			Name:   "resize with fill",
			Image:  "01.jpg",
			Params: []string{"m_fill", "w_100", "h_200"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 100)
				assert.Equal(t, ref.Height(), 200)
			},
		},
		{
			Name:   "resize with pad",
			Image:  "01.jpg",
			Params: []string{"m_pad", "w_100", "h_200"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 100)
				assert.Equal(t, ref.Height(), 200)
			},
		},
		{
			Name:   "resize with fixed",
			Image:  "01.jpg",
			Params: []string{"m_fixed", "w_50", "h_50"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 50)
				assert.Equal(t, ref.Height(), 50)
			},
		},
		{
			Name:   "resize with limit",
			Image:  "01.jpg",
			Params: []string{"m_fixed", "w_800", "h_500", "limit_1"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.False(t, ref.Width() == 800 || ref.Height() == 500)
			},
		},
		{
			Name:   "resize with limit disable",
			Image:  "01.jpg",
			Params: []string{"m_fixed", "w_800", "h_500", "limit_0"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.True(t, ref.Width() == 800 || ref.Height() == 500)
			},
		},
	}

	runTableTest(cases, t, new(Resize))
}
