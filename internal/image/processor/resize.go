package processor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Resize string

func (*Resize) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		m     = "lfit"
		w, h  int
		limit = 1

		// default size
		resizeMode = vips.SizeForce
	)

	log.Printf("resize process with params %#v \n", args.Params)

	for _, param := range args.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, errors.New("invalid resize params")
		}

		switch splits[0] {
		case "m":
			m = splits[1]
		case "w":
			w, err = strconv.Atoi(splits[1])
		case "h":
			h, err = strconv.Atoi(splits[1])
		case "limit":
			limit, err = strconv.Atoi(splits[1])
		}

		if err != nil {
			return nil, err
		}
	}

	// do noting
	if w == 0 && h == 0 {
		return
	}

	imgHeight, imgWidth := args.Img.PageHeight(), args.Img.Width()
	if limit == 1 && (h > imgHeight && w > imgWidth) {
		return nil, nil
	}

	if w == 0 {
		w = h * imgWidth / imgHeight
	} else if h == 0 {
		h = w * imgHeight / imgWidth
	}

	switch m {
	case "lfit":
		if h*imgWidth/imgHeight > imgWidth {
			h = w * imgHeight / imgWidth
		} else {
			w = h * imgWidth / imgHeight
		}
	case "mfit":
		if h*imgWidth/imgHeight > imgWidth {
			w = h * imgWidth / imgHeight
		} else {
			h = w * imgHeight / imgWidth
		}
	case "fill", "pad":
		resizeMode = vips.SizeBoth
	}

	log.Printf("resize with m=%s, w=%d, h=%d, resizeMode=%d \n", m, w, h, resizeMode)
	return nil, args.Img.ThumbnailWithSize(w, h, vips.InterestingCentre, resizeMode)
}
