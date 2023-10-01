package processor

import (
	"errors"
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"

	"github.com/imagecodex/imagecloud/internal/image/metadata"
	"github.com/imagecodex/imagecloud/internal/image/processor/types"
)

type Gray string

func (*Gray) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	if len(args.Params) != 1 {
		return nil, errors.New("invalid gray params")
	}

	value, err := strconv.Atoi(args.Params[0])
	if err != nil {
		return nil, err
	}

	if value == 1 {
		err = args.Img.ToColorSpace(vips.InterpretationBW)
	}

	return
}
