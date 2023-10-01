package processor

import (
	"errors"
	"log"
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/imagecodex/imagecloud/internal/image/metadata"
	"github.com/imagecodex/imagecloud/internal/image/processor/types"
)

type Rotate string

func (*Rotate) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	if len(args.Params) != 1 {
		return nil, errors.New("invalid rotate params")
	}

	angle, err := strconv.Atoi(args.Params[0])
	if err != nil {
		return nil, err
	}

	if angle < 0 || angle > 360 {
		return nil, errors.New("rotate angle should be in range [0,360]")
	}

	log.Printf("rotate image with angle %d", angle)

	switch angle {
	case 0:
		// nothing to do
		return nil, nil
	case 90, 180, 270:
		return nil, args.Img.Rotate(vips.Angle(angle / 90))
	default:
		return nil, args.Img.RotateAny(float64(angle), nil)
	}
}
