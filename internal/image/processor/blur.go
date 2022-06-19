package processor

import (
	"errors"
	"strconv"
	"strings"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Blur string

func (*Blur) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var s int

	for _, param := range args.Params {
		splits := strings.Split(param, "_")
		if len(splits) != 2 {
			return nil, errors.New("invalid blur params")
		}

		switch splits[0] {
		case "s":
			s, err = strconv.Atoi(splits[1])
		}

		if err != nil {
			return
		}
	}

	if s < 1 || s > 50 {
		return nil, errors.New("invalid blur value")
	}

	return nil, args.Img.GaussianBlur(float64(s))
}
