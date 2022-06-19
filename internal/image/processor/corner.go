package processor

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type RoundedCorner string

func (*RoundedCorner) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		r int
	)

	for _, param := range args.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, errors.New("invalid rounded-corner params")
		}

		switch splits[0] {
		case "r":
			r, err = strconv.Atoi(splits[1])
		}

		if err != nil {
			return
		}
	}

	if r <= 0 {
		return
	}

	log.Printf("rounded-corner with r=%d", r)

	imgWith := args.Img.Width()
	imgHeight := args.Img.PageHeight()
	corners := fmt.Sprintf(`<svg viewBox="0 0 %d %d">	
		<rect rx="%d" ry="%d" 
		x="0" y="0" width="%d" height="%d" 
		fill="#FFF"/>
	</svg>`, imgWith, imgHeight, r, r, imgWith, imgHeight)

	tmp, err := vips.NewImageFromBuffer([]byte(corners))
	if err != nil {
		return
	}
	defer tmp.Close()

	if err = tmp.ExtractBand(3, 1); err != nil {
		return
	}

	err = args.Img.BandJoin(tmp)
	return
}
