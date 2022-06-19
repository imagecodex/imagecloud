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
		value int
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

	rate := 1 + float64(value)/100
	err = args.Img.Linear1(rate, 0)
	return nil, err
}
