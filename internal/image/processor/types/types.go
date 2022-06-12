package types

import "github.com/davidbyttow/govips/v2/vips"

type CmdArgs struct {
	Img          *vips.ImageRef
	Ep           *vips.ExportParams
	ObjectPrefix string
	Params       []string
}

func NewCmdArgs(img *vips.ImageRef, ep *vips.ExportParams, prefix string) *CmdArgs {
	return &CmdArgs{
		Img:          img,
		Ep:           ep,
		ObjectPrefix: prefix,
	}
}
