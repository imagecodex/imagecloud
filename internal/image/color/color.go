package color

import (
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
)

func Hex2RGB(hex string) (vips.Color, error) {
	var rgb vips.Color
	values, err := strconv.ParseUint(string(hex), 16, 32)

	if err != nil {
		return vips.Color{}, err
	}

	rgb = vips.Color{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
	}

	return rgb, nil
}
