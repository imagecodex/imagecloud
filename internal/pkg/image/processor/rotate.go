package processor

import (
	"errors"
	"log"
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/pkg/image/metadata"
	"github.com/songjiayang/imagecloud/internal/pkg/image/processor/types"
)

type Rotate string

func (*Rotate) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	if len(args.Params) != 1 {
		return nil, errors.New("invalid rotate params")
	}

	value, err := strconv.ParseInt(args.Params[0], 10, 64)
	if err != nil {
		return nil, err
	}

	var angle vips.Angle

	switch value {
	case 0, 360:
		angle = vips.Angle0
	case 90:
		angle = vips.Angle90
	case 180:
		angle = vips.Angle180
	case 270:
		angle = vips.Angle270
	default:
		return nil, errors.New("roate angle support (0, 90, 180, 270) only")
	}

	log.Printf("rotate image with angle %d \n", angle)

	return nil, args.Img.Rotate(angle)
}
