package processor

import (
	"encoding/base64"

	"github.com/davidbyttow/govips/v2/vips"
)

func base64UrlDecodeString(encodedStr string) (string, error) {
	buf, err := base64.URLEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func getRealOffset(imgWidth, imgHeight, x, y int, g string, boxInfo *vips.ImageMetadata) (int, int) {
	boxWidth, boxHeight := 0, 0
	if boxInfo != nil {
		boxWidth = boxInfo.Width
		boxHeight = boxInfo.Height
	}

	switch g {
	case "north":
		x = imgWidth/2 + x - boxWidth/2
	case "ne":
		x = imgWidth - x - boxWidth
	case "west":
		y = imgHeight/2 + y
	case "center":
		y = imgHeight/2 + y - boxHeight/2
		x = imgWidth/2 + x - boxWidth/2
	case "east":
		y = imgHeight/2 + y - boxHeight/2
		x = imgWidth - x - boxWidth
	case "sw":
		y = imgHeight - y - boxHeight
	case "south":
		y = imgHeight - y - boxHeight
		x = imgWidth/2 + x - boxWidth/2
	case "se":
		y = imgHeight - y - boxHeight
		x = imgWidth - x - boxWidth
	}

	return x, y
}

func ensureInRange(min, max, v int) int {
	if v < min {
		v = min
	} else if v > max {
		v = max
	}

	return v
}
