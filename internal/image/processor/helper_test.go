package processor

import (
	"io/ioutil"
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"

	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

func loadImage(pic string) (*vips.ImageRef, error) {
	data, err := getImageData(pic)
	if err != nil {
		return nil, err
	}

	return vips.NewImageFromBuffer(data)
}

func getImageData(pic string) ([]byte, error) {
	return ioutil.ReadFile("../../../pics/" + pic)
}

type TestCase struct {
	Name      string
	Image     string
	Params    []string
	ExpectErr string
	CheckFunc func(*vips.ImageRef, *testing.T)
}

func runTableTest(cases []TestCase, t *testing.T, p Processor) {
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			ref, err := loadImage(tc.Image)
			assert.Nil(t, err)
			defer ref.Close()

			_, err = p.Process(&types.CmdArgs{
				Img:    ref,
				Params: tc.Params,
			})

			// check error
			if tc.ExpectErr == "" {
				assert.Nil(t, err)
			} else {
				assert.Contains(t, err.Error(), tc.ExpectErr)
			}

			if tc.CheckFunc != nil {
				tc.CheckFunc(ref, t)
			}
		})
	}
}
