package processor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"

	"github.com/songjiayang/imagecloud/internal/image/color"
	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Resize string

func (r *Resize) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		m     = "lfit"
		w, h  int
		limit = 1

		// default size
		resizeMode = vips.SizeForce

		// pad params
		padColor interface{}
	)

	log.Printf("resize process with params %v", args.Params)

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
		case "color":
			padColor, err = r.resolveVipsColor(splits[1])
		}

		if err != nil {
			return nil, err
		}
	}

	// do noting
	if w <= 0 && h <= 0 {
		return
	}

	imgHeight, imgWidth := args.Img.PageHeight(), args.Img.Width()
	if limit == 1 && (h > imgHeight && w > imgWidth) {
		return nil, nil
	}

	if w <= 0 {
		w = h * imgWidth / imgHeight
	} else if h <= 0 {
		h = w * imgHeight / imgWidth
	}

	iw, ih := w, h

	switch m {
	case "lfit", "pad":
		if h*imgWidth/imgHeight >= w {
			h = w * imgHeight / imgWidth
		} else {
			w = h * imgWidth / imgHeight
		}
	case "mfit":
		if h*imgWidth/imgHeight >= w {
			w = h * imgWidth / imgHeight
		} else {
			h = w * imgHeight / imgWidth
		}
	case "fill":
		resizeMode = vips.SizeBoth
	}

	log.Printf("resize with m=%s, w=%d, h=%d, resizeMode=%d", m, w, h, resizeMode)

	// resize first
	if err = args.Img.ThumbnailWithSize(w, h, vips.InterestingCentre, resizeMode); err != nil {
		return nil, err
	}

	// nothing need to change
	if iw == w && ih == h {
		return
	}

	if m == "pad" {
		err = r.pad(args, padColor, iw, ih, w, h)
	}
	return
}

func (*Resize) resolveVipsColor(hexColor string) (interface{}, error) {
	switch len(hexColor) {
	case 6:
		return color.Hex2RGB(hexColor)
	case 8:
		return color.Hex2RGBA(hexColor)
	default:
		return vips.Color{255, 255, 255}, nil
	}
}

func (r *Resize) pad(args *types.CmdArgs, padColor interface{}, iw, ih, w, h int) (err error) {
	if padColor == nil {
		padColor = vips.Color{255, 255, 255}
	}

	left := (iw - w) / 2
	top := (ih - h) / 2

	switch v := padColor.(type) {
	case vips.Color:
		err = args.Img.EmbedBackground(left, top, iw, ih, &v)
	case vips.ColorRGBA:
		err = args.Img.EmbedBackgroundRGBA(left, top, iw, ih, &v)
	}

	return err
}
