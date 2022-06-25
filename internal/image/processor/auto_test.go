package processor

import (
	"io/ioutil"
	"testing"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/stretchr/testify/assert"

	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

func TestAutoRotate(t *testing.T) {
	assertion := assert.New(t)
	ref, err := loadImage("f.jpeg")
	assertion.Nil(err)

	oInfo := ref.Metadata()

	p := new(AutoRotate)

	// if auto-orient disable
	{
		p.Process(&types.CmdArgs{
			Img:    ref,
			Params: []string{"0"},
		})

		info := ref.Metadata()
		assertion.Equal(info.Width, oInfo.Width)
		assertion.Equal(info.Height, oInfo.Height)
	}

	// if auto-orient enable
	{
		p.Process(&types.CmdArgs{
			Img:    ref,
			Params: []string{"1"},
		})

		info := ref.Metadata()
		assertion.Equal(info.Width, oInfo.Height)
		assertion.Equal(info.Height, oInfo.Width)
	}

}

func loadImage(pic string) (*vips.ImageRef, error) {
	data, err := ioutil.ReadFile("../../../pics/" + pic)
	if err != nil {
		return nil, err
	}

	return vips.NewImageFromBuffer(data)
}
