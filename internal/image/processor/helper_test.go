package processor

import (
	"bytes"
	"os"
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"

	"github.com/songjiayang/imagecloud/internal/image/loader"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

func loadImage(pic string) (*vips.ImageRef, error) {
	data, err := getImageData(pic)
	if err != nil {
		return nil, err
	}

	return loader.LoadWithReader(bytes.NewBuffer(data))
}

func getImageData(pic string) ([]byte, error) {
	return os.ReadFile("../../../pics/" + pic)
}

type TestCase struct {
	Name        string
	Image       string
	Params      []string
	ExpectErr   string
	ExportCheck bool
	CheckFunc   func(*vips.ImageRef, *testing.T)
}

func runTableTest(cases []TestCase, t *testing.T, p Processor) {
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			ref, err := loadImage(tc.Image)
			assert.Nil(t, err)
			defer ref.Close()

			args := &types.CmdArgs{
				Img: ref,
				Ep: &vips.ExportParams{
					Format:  ref.Format(),
					Quality: 75,
					Speed:   9,
				},
				Params: tc.Params,
			}

			_, err = p.Process(args)

			// check error
			if tc.ExpectErr == "" {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.ExpectErr)
			}

			if tc.CheckFunc != nil {
				// reload ref
				if tc.ExportCheck {
					if data, _, err := ref.Export(args.Ep); err != nil {
						t.Error(err)
					} else {
						ref.Close()

						if ref, err = vips.NewImageFromBuffer(data); err != nil {
							t.Error(err)
						}
					}
				}

				tc.CheckFunc(ref, t)
			}
		})
	}
}
