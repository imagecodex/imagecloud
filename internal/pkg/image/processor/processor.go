package processor

import (
	"fmt"
	"sync"

	"github.com/songjiayang/imagecloud/internal/pkg/image/metadata"
	"github.com/songjiayang/imagecloud/internal/pkg/image/processor/types"
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
			"resize":    new(Resize),
			"watermark": new(Watermark),
			"crop":      new(Crop),
			"format":    new(Format),
			"info":      new(Info),
			"quality":   new(Quality),
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
