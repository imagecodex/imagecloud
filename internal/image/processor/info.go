package processor

import (
	"github.com/davidbyttow/govips/v2/vips"

	"github.com/imagecodex/imagecloud/internal/image/metadata"
	"github.com/imagecodex/imagecloud/internal/image/processor/types"
)

type Info string

func (*Info) Process(args *types.CmdArgs) (*metadata.Info, error) {
	info := args.Img.Metadata()
	pageHeight := args.Img.GetPageHeight()
	return &metadata.Info{
		Format: vips.ImageTypes[info.Format],
		Height: &pageHeight,
		Width:  &info.Width,
		Pages:  &info.Pages,
	}, nil
}
