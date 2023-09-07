package processor

import (
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	cases := []TestCase{
		{
			Name:        "format to jpg",
			Image:       "01.jpg",
			Params:      []string{"jpg"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeJPEG)
			},
		},
		{
			Name:        "format to jpeg",
			Image:       "01.jpg",
			Params:      []string{"jpeg"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeJPEG)
			},
		},
		{
			Name:        "format to webp",
			Image:       "01.jpg",
			Params:      []string{"webp"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeWEBP)
			},
		},
		{
			Name:        "format to png",
			Image:       "01.jpg",
			Params:      []string{"png"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypePNG)
			},
		},
		{
			Name:        "format to gif",
			Image:       "01.jpg",
			Params:      []string{"gif"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeGIF)
			},
		},
		{
			Name:        "format to avif",
			Image:       "01.jpg",
			Params:      []string{"avif"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeAVIF)
			},
		},
		{
			Name:        "format to heif",
			Image:       "01.jpg",
			Params:      []string{"heif"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeHEIF)
			},
		},
		{
			Name:        "format to heic",
			Image:       "01.jpg",
			Params:      []string{"heic"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeHEIF)
			},
		},
		{
			Name:        "format to jxl",
			Image:       "01.jpg",
			Params:      []string{"jxl"},
			ExportCheck: true,
			CheckFunc: func(ref *vips.ImageRef, t *testing.T) {
				assert.Equal(t, ref.Format(), vips.ImageTypeJXL)
			},
		},
		{
			Name:      "format to invalid type",
			Image:     "01.jpg",
			Params:    []string{"invalid-type"},
			ExpectErr: "image type not support",
		},
	}

	runTableTest(cases, t, new(Format))
}
