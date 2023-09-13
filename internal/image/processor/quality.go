package processor

import (
	"errors"
	"strconv"
	"strings"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Quality string

func (*Quality) Process(args *types.CmdArgs) (*metadata.Info, error) {
	if len(args.Params) != 1 {
		return nil, errors.New("invalid quality params")
	}

LOOP:
	for _, param := range args.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, errors.New("invalid quality params")
		}

		switch splits[0] {
		case "Q":
			quality, err := strconv.ParseInt(splits[1], 10, 64)
			if err != nil {
				return nil, err
			}

			args.Ep.Quality = int(quality)
			break LOOP
		}
	}

	return nil, nil
}
