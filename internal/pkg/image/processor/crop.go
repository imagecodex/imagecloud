package processor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/songjiayang/imagecloud/internal/pkg/image/metadata"
	"github.com/songjiayang/imagecloud/internal/pkg/image/processor/types"
)

type Crop string

func (*Crop) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		w, h, x, y int
		g          string
	)

	for _, param := range args.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, errors.New("invalid resize params")
		}

		switch splits[0] {
		case "w":
			w, err = strconv.Atoi(splits[1])
		case "h":
			h, err = strconv.Atoi(splits[1])
		case "x":
			x, err = strconv.Atoi(splits[1])
		case "y":
			y, err = strconv.Atoi(splits[1])
		case "g":
			g = splits[1]
		}

		if err != nil {
			return
		}
	}

	if w == 0 || h == 0 {
		return nil, errors.New("invalid w, h params")
	}

	metadata := args.Img.Metadata()
	imgHeight, imgWidth := metadata.Height, metadata.Width
	x, y = getRealOffset(imgWidth, imgHeight, x, y, g, nil)
	x = ensureInRange(0, imgWidth, x)
	y = ensureInRange(0, imgHeight, y)

	if x+w > imgWidth {
		w = imgWidth - x
	}

	if y+h > imgHeight {
		h = imgHeight - y
	}

	log.Printf("crop with x=%d, y=%d, w=%d, h=%d", x, y, w, h)

	return nil, args.Img.Crop(x, y, w, h)
}
