package processor

import (
	"fmt"
	"log"

	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type AverageHue string

func (*AverageHue) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	if err = args.Img.Stats(); err != nil {
		log.Printf("image stats with error: %v", err)
		return nil, err
	}

	f1, _ := args.Img.GetPoint(4, 1)
	f2, _ := args.Img.GetPoint(4, 2)
	f3, _ := args.Img.GetPoint(4, 3)

	return &metadata.Info{
		RGB: fmt.Sprintf("0x%x%x%x", int(f1[0]), int(f2[0]), int(f3[0])),
	}, err
}
