package processor

import (
	"errors"
	"fmt"

	"github.com/davidbyttow/govips/v2/vips"

	"github.com/imagecodex/imagecloud/internal/image/metadata"
	"github.com/imagecodex/imagecloud/internal/image/processor/types"
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
	case "avif":
		format = vips.ImageTypeAVIF
	case "heif", "heic":
		format = vips.ImageTypeHEIF
	case "jxl":
		format = vips.ImageTypeJXL
	default:
		return nil, fmt.Errorf("%s image type not support", args.Params[0])
	}

	args.Ep.Format = format

	return nil, nil
}
