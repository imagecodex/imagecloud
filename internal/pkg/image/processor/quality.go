package processor

import (
	"errors"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/pkg/image/metadata"
)

type Quality struct {
	Params []string
}

func NewQuality(params []string) Processor {
	return &Quality{
		Params: params,
	}
}

func (_ *Quality) Name() string {
	return "format"
}

func (q *Quality) Process(_ *vips.ImageRef, ep *vips.ExportParams) (*vips.ImageRef, *metadata.Info, error) {
	if len(q.Params) != 1 {
		return nil, nil, errors.New("invalid quality params")
	}

	for _, param := range q.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, nil, errors.New("invalid quality params")
		}

		switch splits[0] {
		case "Q":
			quality, err := strconv.ParseInt(splits[1], 10, 64)
			if err != nil {
				return nil, nil, err
			}

			ep.Quality = int(quality)
			break
		}
	}

	return nil, nil, nil
}
