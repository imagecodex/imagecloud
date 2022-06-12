package processor

import (
	"errors"
	"strconv"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Bright string

func (*Bright) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		value = 0
	)

	if len(args.Params) != 1 {
		return nil, errors.New("invalid bright params")
	}

	if value, err = strconv.Atoi(args.Params[0]); err != nil {
		return
	}

	if value > 100 || value < -100 {
		return nil, errors.New("invalid bright value, should in range [-100, 100]")
	}

	br := 1 + float64(value)/100

	args.Img.Linear1(br, 0)

	return nil, nil
}
