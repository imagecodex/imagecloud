package processor

import (
	"errors"
	"strconv"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Sharpen string

func (*Sharpen) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		value = 0
	)

	if len(args.Params) != 1 {
		return nil, errors.New("invalid sharpen params")
	}

	if value, err = strconv.Atoi(args.Params[0]); err != nil {
		return
	}

	if value > 399 || value < 50 {
		return nil, errors.New("invalid sharpen value, should in range [50, 399]")
	}

	err = args.Img.Sharpen(0.5, 2, float64(value))
	return nil, err
}
