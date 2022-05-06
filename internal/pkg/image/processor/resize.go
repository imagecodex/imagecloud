package processor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/pkg/image/metadata"
)

type Resize struct {
	Params []string
}

func NewResize(params []string) Processor {
	return &Resize{
		Params: params,
	}
}

func (*Resize) Name() string {
	return "resize"
}

func (r *Resize) Process(img *vips.ImageRef, _ *vips.ExportParams) (*vips.ImageRef, *metadata.Info, error) {
	var (
		m     string
		w, h  int
		limit = 1

		// default size
		resizeMode = vips.SizeForce
	)

	log.Printf("resize process with params %#v \n", r.Params)

	var err error

	for _, param := range r.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, nil, errors.New("invalid resize params")
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
			return nil, nil, err
		}
	}

	imgHeight, imgWidth := img.PageHeight(), img.Width()
	if limit == 1 && (imgHeight > h || imgWidth > w) {
		return img, nil, nil
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
	return img, nil, img.ThumbnailWithSize(w, h, vips.InterestingCentre, resizeMode)
}
