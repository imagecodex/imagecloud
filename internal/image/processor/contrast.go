package processor

import (
	"errors"
	"log"
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Contrast string

func (*Contrast) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		value int
	)

	if len(args.Params) != 1 {
		return nil, errors.New("invalid contrast params")
	}

	if value, err = strconv.Atoi(args.Params[0]); err != nil {
		return
	}

	if value > 100 || value < -100 {
		return nil, errors.New("invalid contrast value, should in range [-100, 100]")
	}

	rate := 1 + float64(value)/100

	if err = args.Img.ToColorSpace(vips.InterpretationLAB); err != nil {
		log.Printf("covert image to lab with error: %v", err)
		return
	}

	if err = args.Img.Linear([]float64{1, rate, rate}, []float64{0, 0, 0}); err != nil {
		log.Printf("image linear with error: %v", err)
		return
	}

	if err = args.Img.ToColorSpace(vips.InterpretationSRGB); err != nil {
		log.Printf("covert image to srgb with error: %v", err)
		return
	}

	return nil, err
}
