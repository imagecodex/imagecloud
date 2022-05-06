package processor

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/pkg/image/metadata"
)

type Info struct{}

func (*Info) Name() string {
	return "info"
}

func (*Info) Process(img *vips.ImageRef, _ *vips.ExportParams) (*vips.ImageRef, *metadata.Info, error) {
	info := img.Metadata()
	return img, &metadata.Info{
		Format: vips.ImageTypes[info.Format],
		Height: info.Height,
		Width:  info.Width,
		Pages:  info.Pages,
	}, nil
}
