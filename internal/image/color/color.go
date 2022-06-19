package color

import (
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
)

func Hex2RGB(hex string) (vips.Color, error) {
	bands, err := strconv.ParseUint(string(hex), 16, 32)
	if err != nil {
		return vips.Color{}, err
	}

	rgb := vips.Color{
		R: uint8(bands >> 16),
		G: uint8((bands >> 8) & 0xFF),
		B: uint8(bands & 0xFF),
	}

	return rgb, nil
}
