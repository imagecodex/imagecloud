package processor

import (
	"errors"
	"fmt"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imgcloud/internal/pkg/image/metadata"
)

type Format struct {
	Params []string
}

func NewFormat(params []string) Processor {
	return &Format{
		Params: params,
	}
}

func (_ *Format) Name() string {
	return "format"
}

func (f *Format) Process(_ *vips.ImageRef, ep *vips.ExportParams) (*vips.ImageRef, *metadata.Info, error) {
	if len(f.Params) != 1 {
		return nil, nil, errors.New("invalid format params")
	}

	switch f.Params[0] {
	case "jpg", "jpeg":
		ep.Format = vips.ImageTypeJPEG
	case "webp":
		ep.Format = vips.ImageTypeWEBP
	case "png":
		ep.Format = vips.ImageTypePNG
	case "gif":
		ep.Format = vips.ImageTypeGIF
	default:
		return nil, nil, fmt.Errorf("%s image type not support", f.Params[0])
	}

	return nil, nil, nil
}
