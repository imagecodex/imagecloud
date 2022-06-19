package processor

import (
	"errors"
	"strconv"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type AutoRotate string

func (*AutoRotate) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	if len(args.Params) != 1 {
		return nil, errors.New("invalid auto orient params")
	}

	value, err := strconv.ParseInt(args.Params[0], 10, 64)
	if err != nil {
		return nil, err
	}

	switch value {
	case 0, 1:
	default:
		return nil, errors.New("auto orient value support (0, 1) only")
	}

	// ignore if value != 1
	if value != 1 {
		return
	}

	return nil, args.Img.AutoRotate()
}
