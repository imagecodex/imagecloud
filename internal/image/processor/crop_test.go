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
		{
			Name:   "crop gif with g_nw",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_nw"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_ne",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_ne"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_west",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_west"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_center",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_center"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_east",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_east"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_sw",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_sw"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_south",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_south"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_north",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_north"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
		{
			Name:   "crop gif with g_se",
			Image:  "01.gif",
			Params: []string{"x_20", "y_20", "w_200", "h_200", "g_se"},
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Width(), 200)
				assert.Equal(t, ref.PageHeight(), 200)
				assert.Equal(t, ref.Pages(), 3)
			},
		},
	}

	runTableTest(cases, t, new(Crop))
}
