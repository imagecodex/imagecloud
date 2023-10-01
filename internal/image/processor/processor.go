package processor

import (
	"fmt"
	"sync"

	"github.com/imagecodex/imagecloud/internal/image/metadata"
	"github.com/imagecodex/imagecloud/internal/image/processor/types"
)

type Processor interface {
	Process(*types.CmdArgs) (*metadata.Info, error)
}

var (
	pMaps map[string]Processor
	once  sync.Once
)

func init() {
	once.Do(func() {
		pMaps = map[string]Processor{
			"resize":          new(Resize),
			"watermark":       new(Watermark),
			"crop":            new(Crop),
			"format":          new(Format),
			"info":            new(Info),
			"auto-orient":     new(AutoRotate),
			"circle":          new(Circle),
			"indexcrop":       new(IndexCrop),
			"rounded-corners": new(RoundedCorner),
			"quality":         new(Quality),
			"blur":            new(Blur),
			"rotate":          new(Rotate),
			"average-hue":     new(AverageHue),
			"bright":          new(Bright),
			"sharpen":         new(Sharpen),
			"contrast":        new(Contrast),
			"gray":            new(Gray),
		}
	})
}

func Excute(name string, args *types.CmdArgs) (*metadata.Info, error) {
	if p := pMaps[name]; p == nil {
		return nil, fmt.Errorf("no prossor for command %s", name)
	} else {
		return p.Process(args)
	}
}
