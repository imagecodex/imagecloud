package processor

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imgcloud/internal/pkg/image/metadata"
)

type Processor interface {
	Name() string
	Process(*vips.ImageRef, *vips.ExportParams) (*vips.ImageRef, *metadata.Info, error)
}

type ProcessorNewFunc func(params []string) Processor

var ProcessorNewMap = map[string]ProcessorNewFunc{
	"resize":  NewResize,
	"format":  NewFormat,
	"quality": NewQuality,
	"info":    func([]string) Processor { return &Info{} },
}

func registerFunc(name string, pf ProcessorNewFunc) {
	ProcessorNewMap[name] = pf
}
