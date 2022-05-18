package processor

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Info string

func (*Info) Process(args *types.CmdArgs) (*metadata.Info, error) {
	info := args.Img.Metadata()
	return &metadata.Info{
		Format: vips.ImageTypes[info.Format],
		Height: info.Height,
		Width:  info.Width,
		Pages:  info.Pages,
	}, nil
}
