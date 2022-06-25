package processor

import (
	"testing"

	"github.com/songjiayang/imagecloud/internal/image/processor/types"
	"github.com/stretchr/testify/assert"
)

func TestBlur(t *testing.T) {
	assertion := assert.New(t)

	ref, err := loadImage("01.jpg")
	assertion.Nil(err)
	defer ref.Close()

	b := new(Blur)

	{
		_, err = b.Process(&types.CmdArgs{
			Img:    ref,
			Params: []string{"s"},
		})
		assertion.Contains(err.Error(), "invalid blur params")
	}

	{
		_, err = b.Process(&types.CmdArgs{
			Img:    ref,
			Params: []string{"s_100"},
		})
		assertion.Contains(err.Error(), "invalid blur value")
	}

	{
		_, err = b.Process(&types.CmdArgs{
			Img:    ref,
			Params: []string{"s_2"},
		})
		assertion.Nil(err)
	}
}
