package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex2RGB(t *testing.T) {
	assertion := assert.New(t)

	{
		_, err := Hex2RGB("7ad42x")
		assertion.NotNil(err)
	}

	{
		color, err := Hex2RGB("7ad42f")
		assertion.Nil(err)
		assertion.Equal(color.R, uint8(0x7a))
		assertion.Equal(color.G, uint8(0xd4))
		assertion.Equal(color.B, uint8(0x2f))
	}
}
