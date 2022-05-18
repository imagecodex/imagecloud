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

	var auto bool

	switch value {
	case 0:
		auto = false
	case 1:
		auto = true
	default:
		return nil, errors.New("auto orient value support (0, 1) only")
	}

	if !auto {
		return
	}

	return nil, args.Img.AutoRotate()
}
