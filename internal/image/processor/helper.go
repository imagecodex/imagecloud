package processor

import (
	"encoding/base64"

	"github.com/davidbyttow/govips/v2/vips"
)

func base64UrlDecodeString(encodedStr string) (string, error) {
	if encodedStr == "" {
		return "", nil
	}

	buf, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		buf, err = base64.URLEncoding.DecodeString(encodedStr)
	}
	if err != nil {
		buf, err = base64.RawStdEncoding.DecodeString(encodedStr)
	}
	if err != nil {
		buf, err = base64.RawURLEncoding.DecodeString(encodedStr)
	}

	return string(buf), err
}

func getRealOffset(imgWidth, imgHeight, x, y int, g string, box *vips.ImageMetadata) (int, int) {
	boxWidth, boxHeight := 0, 0
	if box != nil {
		boxWidth = box.Width
		boxHeight = box.Height
	}

	switch g {
	case "north":
		x = imgWidth/2 + x - boxWidth/2
	case "ne":
		x = imgWidth - x - boxWidth
	case "west":
		y = imgHeight/2 - boxHeight/2 + y
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
	case "nw":
		// do noting
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
