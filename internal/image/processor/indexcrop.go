package processor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type IndexCrop string

func (*IndexCrop) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		x, y, i int
	)

	for _, param := range args.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, errors.New("invalid indexcrop params")
		}

		switch splits[0] {
		case "x":
			x, err = strconv.Atoi(splits[1])
		case "y":
			y, err = strconv.Atoi(splits[1])
		case "i":
			i, err = strconv.Atoi(splits[1])
		}

		if err != nil {
			return
		}
	}

	imgWidth, imgHeight := args.Img.Width(), args.Img.PageHeight()
	// ignore if x, y == 0
	if x <= 0 && y <= 0 {
		log.Printf("ignore indexcrop x,y both zero")
		return
	}

	if i < 0 {
		i = 0
	}

	var left, top, w, h int
	if x > 0 {
		left = x * i
		w = x
		h = imgHeight
	} else {
		top = y * i
		w = imgWidth
		h = y
	}

	// if the top,left out of the image box
	if left > imgWidth || top > imgHeight {
		return
	}

	// if crop area output image, update the crop w, h
	if left+w > imgWidth {
		w = imgWidth - left
	}
	if top+h > imgHeight {
		h = imgHeight - top
	}

	log.Printf("crop with left=%d, top=%d, w=%d, h=%d", left, top, w, h)
	err = args.Img.Crop(left, top, w, h)
	return
}
