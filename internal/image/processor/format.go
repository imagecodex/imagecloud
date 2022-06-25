package processor

import (
	"errors"
	"fmt"

	"github.com/davidbyttow/govips/v2/vips"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Format string

func (*Format) Process(args *types.CmdArgs) (*metadata.Info, error) {
	if len(args.Params) != 1 {
		return nil, errors.New("invalid format params")
	}

	var format vips.ImageType

	switch args.Params[0] {
	case "jpg", "jpeg":
		format = vips.ImageTypeJPEG
	case "webp":
		format = vips.ImageTypeWEBP
	case "png":
		format = vips.ImageTypePNG
	case "gif":
		format = vips.ImageTypeGIF
	default:
		return nil, fmt.Errorf("%s image type not support", args.Params[0])
	}

	args.Ep.Format = format

	return nil, nil
}
