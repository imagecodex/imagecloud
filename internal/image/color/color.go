package color

import (
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
)

func Hex2RGB(hex string) (c vips.Color, err error) {
	var bands uint64
	if bands, err = strconv.ParseUint(string(hex), 16, 64); err != nil {
		return
	}

	c = vips.Color{
		R: uint8(bands >> 16),
		G: uint8((bands >> 8) & 0xFF),
		B: uint8(bands & 0xFF),
	}
	return
}

func Hex2RGBA(hex string) (c vips.ColorRGBA, err error) {
	var bands uint64
	if bands, err = strconv.ParseUint(string(hex), 16, 64); err != nil {
		return
	}

	c = vips.ColorRGBA{
		R: uint8(bands >> 24),
		G: uint8((bands >> 16) & 0xFF),
		B: uint8((bands >> 8) & 0xFF),
		A: uint8(bands & 0xFF),
	}
	return
}
